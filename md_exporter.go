package acmelib

import (
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	md "github.com/nao1215/markdown"
	"golang.org/x/exp/maps"
)

// ExportToMarkdown exports the given [Network] to a markdown document.
// It writes the markdown document to the given [io.Writer].
func ExportToMarkdown(network *Network, w io.Writer) error {
	mdWriter := md.NewMarkdown(w)
	exporter := newMDExporter(mdWriter)
	exporter.exportNetwork(network)
	return mdWriter.Build()
}

type mdExporter struct {
	w *md.Markdown

	sigTypes map[EntityID]*SignalType
	sigUnits map[EntityID]*SignalUnit
	sigEnums map[EntityID]*SignalEnum
}

func newMDExporter(mdWriter *md.Markdown) *mdExporter {
	return &mdExporter{
		w: mdWriter,

		sigTypes: make(map[EntityID]*SignalType),
		sigUnits: make(map[EntityID]*SignalUnit),
		sigEnums: make(map[EntityID]*SignalEnum),
	}
}

func (e *mdExporter) getHeaderLink(headerName string) string {
	linkStr := "#" + strings.ReplaceAll(strings.ToLower(headerName), " ", "-")
	linkStr = strings.ReplaceAll(linkStr, "`", "")
	return md.Link(headerName, linkStr)
}

func (e *mdExporter) exportNetwork(net *Network) {
	e.w.Importantf("This markdown document is generated by %s", md.Link("acmelib", "https://github.com/squadracorsepolito/acmelib")).LF()

	e.w.H1(net.name)

	if len(net.desc) > 0 {
		e.w.PlainText(net.desc).LF()
		e.w.HorizontalRule()
	}

	e.exportTOC(net)

	for _, bus := range net.Buses() {
		e.exportBus(bus)
	}

	sigTypes := maps.Values(e.sigTypes)
	slices.SortFunc(sigTypes, func(a, b *SignalType) int { return a.size - b.size })
	e.exportSignalTypes(sigTypes)

	sigUnits := maps.Values(e.sigUnits)
	slices.SortFunc(sigUnits, func(a, b *SignalUnit) int { return strings.Compare(a.name, b.name) })
	e.exportSignalUnits(sigUnits)

	sigEnums := maps.Values(e.sigEnums)
	slices.SortFunc(sigEnums, func(a, b *SignalEnum) int { return strings.Compare(a.name, b.name) })
	e.w.H2("Signal Enums")
	e.w.PlainText("The list of all the signal enums used in the network.").LF()
	for _, tmpSigEnum := range sigEnums {
		e.exportSignalEnum(tmpSigEnum)
	}
}

func (e *mdExporter) exportTOC(net *Network) {
	for _, bus := range net.Buses() {
		e.w.BulletList(e.getHeaderLink(bus.name))
		for _, nodeInt := range bus.NodeInterfaces() {
			e.w.PlainTextf("\t- %s", e.getHeaderLink(nodeInt.node.name))
			for _, msg := range nodeInt.Messages() {
				e.w.PlainTextf("\t\t- %s", e.getHeaderLink(msg.name))
			}
		}
	}

	e.w.BulletList(e.getHeaderLink("Signal Types"))
	e.w.BulletList(e.getHeaderLink("Signal Units"))
	e.w.BulletList(e.getHeaderLink("Signal Enums"))
}

func (e *mdExporter) exportBus(bus *Bus) {
	e.w.H2(bus.name)

	if len(bus.desc) > 0 {
		e.w.PlainText(bus.desc).LF()
		e.w.HorizontalRule()
	}

	baudrateStr := "-"
	if bus.baudrate != 0 {
		baudrateStr = md.Bold(strconv.Itoa(bus.baudrate))
	}
	e.w.PlainTextf("Baudrate: %s bps", baudrateStr).LF()

	for _, node := range bus.NodeInterfaces() {
		e.exportNode(node)
	}
}

func (e *mdExporter) exportNode(node *NodeInterface) {
	e.w.HorizontalRule()
	e.w.H3(node.node.name)

	if len(node.node.desc) > 0 {
		e.w.PlainText(node.node.desc).LF()
		e.w.HorizontalRule()
	}

	e.w.PlainTextf("Node ID: %s", md.Bold(fmt.Sprintf("%d", node.node.id))).LF()

	for _, msg := range node.Messages() {
		e.exportMessage(msg)
	}
}

func (e *mdExporter) exportMessage(msg *Message) {
	e.w.HorizontalRule()
	e.w.H4(msg.name)

	if len(msg.desc) > 0 {
		e.w.PlainText(msg.desc).LF()
		e.w.HorizontalRule()
	}

	if msg.hasStaticCANID {
		e.w.PlainTextf("CAN-ID: %s (static)", md.Bold(fmt.Sprintf("%d", msg.staticCANID))).LF()
	} else {
		e.w.PlainTextf("Message ID: %s", md.Bold(fmt.Sprintf("%d", msg.id))).LF()
		e.w.PlainTextf("CAN-ID: %s (generated)", md.Bold(fmt.Sprintf("%d", msg.GetCANID()))).LF()
	}
	e.w.PlainTextf("Size: %s bytes", md.Bold(fmt.Sprintf("%d", msg.sizeByte))).LF()
	e.w.PlainTextf("Byte Order: %s", md.Bold(msg.byteOrder.String())).LF()

	cycleTimeStr := "-"
	if msg.cycleTime > 0 {
		cycleTimeStr = fmt.Sprintf("%s ms", md.Bold(strconv.Itoa(msg.cycleTime)))
	}
	e.w.PlainTextf("Cycle Time: %s", cycleTimeStr).LF()

	recStr := "Receivers: "
	for idx, recInt := range msg.Receivers() {
		recLink := e.getHeaderLink(recInt.node.name)

		if idx == 0 {
			recStr += recLink
			continue
		}

		recStr += fmt.Sprintf(", %s", recLink)
	}
	e.w.PlainText(recStr).LF()

	if len(msg.Signals()) == 0 {
		return
	}

	sigTable := md.TableSet{
		Header: []string{"Name", "Start Bit", "Size", "Type", "Min", "Max", "Unit", "Description"},
		Rows:   [][]string{},
	}
	for _, sig := range msg.Signals() {
		sigRows := e.exportSignal(sig)
		for _, tmpSigRow := range sigRows {
			sigTable.Rows = append(sigTable.Rows, tmpSigRow)
		}
	}
	e.w.CustomTable(sigTable, md.TableOptions{AutoWrapText: false, AutoFormatHeaders: false})
}

func (e *mdExporter) exportSignal(sig Signal) (resRows [][]string) {
	sigRow := []string{}

	sigRow = append(sigRow, sig.Name())
	sigRow = append(sigRow, fmt.Sprintf("%d", sig.GetStartBit()))
	sigRow = append(sigRow, fmt.Sprintf("%d", sig.GetSize()))

	switch sig.Kind() {
	case SignalKindStandard:
		stdSig, err := sig.ToStandard()
		if err != nil {
			panic(err)
		}
		sigRow = append(sigRow, e.exportStandardSignal(stdSig)...)
		resRows = append(resRows, sigRow)

	case SignalKindEnum:
		enumSig, err := sig.ToEnum()
		if err != nil {
			panic(err)
		}
		sigRow = append(sigRow, e.exportEnumSignal(enumSig)...)
		resRows = append(resRows, sigRow)

	case SignalKindMultiplexer:
		muxSig, err := sig.ToMultiplexer()
		if err != nil {
			panic(err)
		}
		muxSigRows := e.exportMultiplexerSignal(muxSig)
		for _, tmpRow := range muxSigRows {
			resRows = append(resRows, tmpRow)
		}
	}

	return resRows
}

func (e *mdExporter) exportStandardSignal(stdSig *StandardSignal) (resRow []string) {
	sigType := stdSig.typ
	e.sigTypes[sigType.entityID] = sigType
	resRow = append(resRow, md.Link(md.Code(sigType.name), "#signal-types"))

	resRow = append(resRow, fmt.Sprintf("%g", stdSig.typ.min))
	resRow = append(resRow, fmt.Sprintf("%g", stdSig.typ.max))

	unitSymbol := "-"
	if stdSig.unit != nil {
		sigUnit := stdSig.unit
		unitSymbol = md.Link(md.Code(sigUnit.symbol), "#signal-units")
		e.sigUnits[sigUnit.entityID] = sigUnit
	}
	resRow = append(resRow, unitSymbol)

	desc := stdSig.desc
	if len(desc) == 0 {
		desc = "-"
	}
	resRow = append(resRow, desc)

	return resRow
}

func (e *mdExporter) exportEnumSignal(enumSig *EnumSignal) (resRow []string) {
	sigEnum := enumSig.enum
	e.sigEnums[sigEnum.entityID] = sigEnum
	resRow = append(resRow, e.getHeaderLink(sigEnum.name))

	resRow = append(resRow, "0")
	resRow = append(resRow, fmt.Sprintf("%d", sigEnum.maxIndex))
	resRow = append(resRow, "-")

	desc := enumSig.desc
	if len(desc) == 0 {
		desc = "-"
	}
	resRow = append(resRow, desc)

	return resRow
}

func (e *mdExporter) exportMultiplexerSignal(muxSig *MultiplexerSignal) (resRows [][]string) {
	muxSigRow := []string{}
	muxSigRow = append(muxSigRow, md.Code("multiplexer"))

	muxSigRow = append(muxSigRow, "0")
	muxSigRow = append(muxSigRow, fmt.Sprintf("%d", muxSig.groupCount))
	muxSigRow = append(muxSigRow, "-")

	desc := muxSig.desc
	if len(desc) == 0 {
		desc = "-"
	}
	muxSigRow = append(muxSigRow, desc)

	resRows = append(resRows, muxSigRow)

	for groupID, group := range muxSig.GetSignalGroups() {
		tmpCol := fmt.Sprintf("- %d -", groupID)
		resRows = append(resRows, []string{tmpCol, tmpCol, tmpCol, tmpCol, tmpCol, tmpCol, tmpCol, tmpCol})

		for _, tmpSig := range group {
			sigRows := e.exportSignal(tmpSig)
			for _, tmpSigRow := range sigRows {
				resRows = append(resRows, tmpSigRow)
			}
		}
	}

	return resRows
}

func (e *mdExporter) exportSignalTypes(sigTypes []*SignalType) {
	e.w.H2("Signal Types")
	e.w.PlainText("The list of all the signal types used in the network.").LF()

	sigTypeTable := md.TableSet{
		Header: []string{"Name", "Size", "Kind", "Signed", "Min", "Max", "Scale", "Offset", "Description"},
		Rows:   [][]string{},
	}
	for _, tmpSigType := range sigTypes {
		desc := "-"
		if len(tmpSigType.desc) > 0 {
			desc = tmpSigType.desc
		}

		sizeStr := fmt.Sprintf("%d", tmpSigType.size)
		signedStr := fmt.Sprintf("%t", tmpSigType.signed)
		minStr := fmt.Sprintf("%g", tmpSigType.min)
		maxStr := fmt.Sprintf("%g", tmpSigType.max)
		scaleStr := fmt.Sprintf("%g", tmpSigType.scale)
		offsetStr := fmt.Sprintf("%g", tmpSigType.offset)

		row := []string{tmpSigType.name, sizeStr, md.Code(tmpSigType.kind.String()), md.Code(signedStr), minStr, maxStr, scaleStr, offsetStr, desc}
		sigTypeTable.Rows = append(sigTypeTable.Rows, row)
	}
	e.w.CustomTable(sigTypeTable, md.TableOptions{AutoWrapText: false, AutoFormatHeaders: false})
}

func (e *mdExporter) exportSignalUnits(sigUnits []*SignalUnit) {
	e.w.H2("Signal Units")
	e.w.PlainText("The list of all the signal units used in the network.").LF()

	sigUnitTable := md.TableSet{
		Header: []string{"Name", "Kind", "Symbol", "Description"},
		Rows:   [][]string{},
	}
	for _, tmpSigUnit := range sigUnits {
		desc := "-"
		if len(tmpSigUnit.desc) > 0 {
			desc = tmpSigUnit.desc
		}

		row := []string{tmpSigUnit.name, tmpSigUnit.kind.String(), tmpSigUnit.symbol, desc}
		sigUnitTable.Rows = append(sigUnitTable.Rows, row)
	}
	e.w.CustomTable(sigUnitTable, md.TableOptions{AutoWrapText: false, AutoFormatHeaders: false})
}

func (e *mdExporter) exportSignalEnum(sigEnum *SignalEnum) {
	e.w.HorizontalRule()
	e.w.H4(sigEnum.name)

	if len(sigEnum.desc) > 0 {
		e.w.PlainText(sigEnum.desc).LF()
		e.w.HorizontalRule()
	}

	valueTable := md.TableSet{
		Header: []string{"Name", "Index", "Description"},
		Rows:   [][]string{},
	}
	for _, tmpVal := range sigEnum.Values() {
		desc := "-"
		if len(tmpVal.desc) > 0 {
			desc = tmpVal.desc
		}
		row := []string{tmpVal.name, fmt.Sprintf("%d", tmpVal.index), desc}
		valueTable.Rows = append(valueTable.Rows, row)
	}
	e.w.CustomTable(valueTable, md.TableOptions{AutoWrapText: false, AutoFormatHeaders: false})
}
