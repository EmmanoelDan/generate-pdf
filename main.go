package main

import (
	"fmt"
	"os"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"

	"github.com/EmmanoelDan/generate-pdf/data"
)
func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHeading(m)
	buildMusicalInstrumentList(m)

	err := m.OutputFileAndClose("pdfs/div_generate.pdf")

	if err != nil {
		fmt.Println("could not save PDF", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved sucessfully.")
}

func buildHeading(m pdf.Maroto){
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("images/logo_div_generate.png", props.Rect{
					Center: true,
					Percent: 500,
				})

				if err!= nil {
                    fmt.Println("Image file was not loaded -", err)
                }
			})
		})
	})

	m.Row(10, func (){
		m.Col(12, func() {
            m.Text("Lista de Instrumentos - Projeto Musica na comunidade", props.Text{
				Top: 3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getDarkPurpleColor(),
			})
        })
	})
}

func buildMusicalInstrumentList(m pdf.Maroto){
	tableHeadings := []string{"Instrumento", "Responsavel", "Situação"}
	contents := data.MusicalInstrumentList(20)
	lightPurpleColor := getLightPurpleColor();
	m.SetBackgroundColor(
		getTealColor(),
	)

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Instrumentos", props.Text{
				Top: 2,
				Size: 13,
				Color: color.NewWhite(),
				Family: consts.Courier,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size: 9,
			GridSizes: []uint{3, 7, 2} ,
		},
		ContentProp: props.TableListContent{
			Size: 8,
            GridSizes: []uint{3, 7, 2},
		},
		Align: consts.Left,
		HeaderContentSpace: 1,
		Line: false,
		AlternatedBackground: &lightPurpleColor,

	})
}
func getDarkPurpleColor() color.Color {
	return color.Color {
		Red: 88,
        Green: 80,
        Blue: 99,
	}
}

func getLightPurpleColor() color.Color {
	return color.Color {
        Red: 210,
        Green: 200,
        Blue: 230,
    }
}

func getTealColor() color.Color {
	return color.Color {
		Red: 3,
        Green: 166,
        Blue: 166,
	}
}