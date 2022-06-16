package tui

import (
	"context"
	"testing"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/strangelove-ventures/ibctest/internal/blockdb"
	"github.com/stretchr/testify/require"
)

var (
	escKey   = tcell.NewEventKey(tcell.KeyESC, ' ', 0)
	enterKey = tcell.NewEventKey(tcell.KeyEnter, ' ', 0)
)

func runKey(c rune) *tcell.EventKey {
	return tcell.NewEventKey(tcell.KeyRune, c, 0)
}

// draw is necessary for some of the below tests to get default behavior such as selecting the first available
// row in a *tview.Table.
func draw(view tview.Primitive) {
	view.Draw(tcell.NewSimulationScreen(""))
}

type mockQueryService struct {
	GotChainPkey int64
	Messages     []blockdb.CosmosMessageResult
	Txs          []blockdb.TxResult
	Err          error
}

func (m *mockQueryService) Transactions(ctx context.Context, chainPkey int64) ([]blockdb.TxResult, error) {
	if ctx == nil {
		panic("nil context")
	}
	m.GotChainPkey = chainPkey
	return m.Txs, m.Err
}

func (m *mockQueryService) CosmosMessages(ctx context.Context, chainPkey int64) ([]blockdb.CosmosMessageResult, error) {
	if ctx == nil {
		panic("nil context")
	}
	m.GotChainPkey = chainPkey
	return m.Messages, m.Err
}

func TestModel_Update(t *testing.T) {
	ctx := context.Background()

	t.Run("go back", func(t *testing.T) {
		model := NewModel(&mockQueryService{}, "", "", time.Now(), nil)
		require.Equal(t, 1, model.mainContentView().GetPageCount())

		draw(model.RootView())

		update := model.Update(ctx)
		update(escKey)

		require.Equal(t, 1, model.mainContentView().GetPageCount())
	})

	t.Run("cosmos summary view", func(t *testing.T) {
		querySvc := &mockQueryService{
			Messages: []blockdb.CosmosMessageResult{
				{Height: 10},
				{Height: 11},
				{Height: 12},
			},
		}
		model := NewModel(querySvc, "", "", time.Now(), []blockdb.TestCaseResult{
			{ChainPKey: 5, ChainID: "my-chain1"},
			{ChainPKey: 6},
		})

		draw(model.RootView())

		update := model.Update(ctx)
		update(runKey('m'))

		// By default, first row is selected in a rendered table.
		require.EqualValues(t, 5, querySvc.GotChainPkey)

		require.Equal(t, 2, model.mainContentView().GetPageCount())
		_, table := model.mainContentView().GetFrontPage()

		// 4 rows: 1 header + 3 blockdb.CosmosMessageResult
		require.Equal(t, 4, table.(*tview.Table).GetRowCount())
		require.Contains(t, table.(*tview.Table).GetTitle(), "my-chain1")
	})

	t.Run("tx detail view", func(t *testing.T) {
		querySvc := &mockQueryService{
			Txs: []blockdb.TxResult{
				{Height: 12, Tx: []byte(`{"tx":1}`)},
				{Height: 13, Tx: []byte(`{"tx":2}`)},
				{Height: 14, Tx: []byte(`{"tx":3}`)},
			},
		}
		model := NewModel(querySvc, "", "", time.Now(), []blockdb.TestCaseResult{
			{ChainPKey: 5, ChainID: "my-chain1"},
			{ChainPKey: 6},
		})

		draw(model.RootView())

		update := model.Update(ctx)

		update(enterKey)

		// By default, first row is selected in a rendered table.
		require.EqualValues(t, 5, querySvc.GotChainPkey)

		require.Equal(t, 2, model.mainContentView().GetPageCount())
		_, primitive := model.mainContentView().GetFrontPage()
		innerPages := primitive.(*tview.Pages)

		require.Equal(t, 3, innerPages.GetPageCount())

		_, primitive = innerPages.GetFrontPage()
		textView := primitive.(*tview.TextView)

		require.Contains(t, textView.GetTitle(), "Tx 1 of 3")
		require.Contains(t, textView.GetTitle(), "my-chain1 @ Height 12")
		const wantFirstPage = `{
  "tx": 1
}`
		require.Equal(t, wantFirstPage, textView.GetText(true))

		// Move to the next page.
		update(runKey(']'))

		_, primitive = innerPages.GetFrontPage()
		textView = primitive.(*tview.TextView)

		require.Contains(t, textView.GetTitle(), "Tx 2 of 3")
		require.Contains(t, textView.GetTitle(), "my-chain1 @ Height 13")
		const wantSecondPage = `{
  "tx": 2
}`
		require.Equal(t, wantSecondPage, textView.GetText(true))

		// Assert does not advance past last page.
		update(runKey(']'))
		update(runKey(']'))
		update(runKey(']'))

		_, primitive = innerPages.GetFrontPage()
		textView = primitive.(*tview.TextView)

		require.Contains(t, textView.GetTitle(), "Tx 3 of 3")

		// Move back to the previous page. Assert does not retreat past first page.
		update(runKey('['))
		update(runKey('['))
		update(runKey('['))
		update(runKey('['))

		_, primitive = innerPages.GetFrontPage()
		textView = primitive.(*tview.TextView)

		require.Contains(t, textView.GetTitle(), "Tx 1 of 3")
	})
}