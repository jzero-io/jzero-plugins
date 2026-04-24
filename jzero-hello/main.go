/*
Copyright © 2026 jaronnie <jaron@jaronnie.com>

A simple example jzero plugin that demonstrates the plugin system.
*/

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jzero-io/jzero/cmd/jzero/pkg/plugin"
	"github.com/spf13/cobra"
	ddlparser "github.com/zeromicro/ddl-parser/parser"
)

var rootCmd = &cobra.Command{
	Use:   "jzero-hello",
	Short: "A simple hello plugin for jzero",
	Long: `This is an example plugin for jzero that demonstrates
how to create and structure a jzero plugin.

Plugins must be named with the "jzero-" prefix.`,
}

var descCmd = &cobra.Command{
	Use:   "desc",
	Short: "Show plugin descriptor information",
	Long:  `Display information about API, Proto, and Model specifications found in the plugin.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := plugin.New()
		if err != nil {
			return err
		}

		if p.Desc.Api.SpecMap != nil {
			for file, spec := range p.Desc.Api.SpecMap {
				fmt.Printf("api file: %s\n", file)
				for _, group := range spec.Service.Groups {
					fmt.Printf("  group: %s\n", group.GetAnnotation("group"))
					for _, route := range group.Routes {
						fmt.Printf("    %s\n", fmt.Sprintf("%s:%s%s", route.Method, group.GetAnnotation("prefix"), route.Path))
					}
				}
			}
		}

		if p.Desc.Proto.SpecMap != nil {
			for file, spec := range p.Desc.Proto.SpecMap {
				fmt.Printf("proto file: %s\n", file)
				for _, service := range spec.Service {
					for _, rpc := range service.RPC {
						fmt.Printf("  %s.%s\n", service.Name, rpc.Name)
					}
				}
			}
		}

		if p.Desc.Model.SpecMap != nil {
			for file, spec := range p.Desc.Model.SpecMap {
				fmt.Printf("sql file: %s\n", file)
				for _, columns := range spec.Columns {
					fmt.Printf("  %s(%s)\n", columns.Name, formatDataType(columns.DataType))
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(descCmd)
}

var dataTypeNames = map[int]string{
	ddlparser.LongVarBinary:      "LongVarBinary",
	ddlparser.LongVarChar:        "LongVarChar",
	ddlparser.GeometryCollection: "GeometryCollection",
	ddlparser.GeomCollection:     "GeomCollection",
	ddlparser.LineString:         "LineString",
	ddlparser.MultiLineString:    "MultiLineString",
	ddlparser.MultiPoint:         "MultiPoint",
	ddlparser.MultiPolygon:       "MultiPolygon",
	ddlparser.Point:              "Point",
	ddlparser.Polygon:            "Polygon",
	ddlparser.Json:               "Json",
	ddlparser.Geometry:           "Geometry",
	ddlparser.Enum:               "Enum",
	ddlparser.Set:                "Set",
	ddlparser.Bit:                "Bit",
	ddlparser.Time:               "Time",
	ddlparser.Timestamp:          "Timestamp",
	ddlparser.DateTime:           "DateTime",
	ddlparser.Binary:             "Binary",
	ddlparser.VarBinary:          "VarBinary",
	ddlparser.Blob:               "Blob",
	ddlparser.Year:               "Year",
	ddlparser.Decimal:            "Decimal",
	ddlparser.Dec:                "Dec",
	ddlparser.Fixed:              "Fixed",
	ddlparser.Numeric:            "Numeric",
	ddlparser.Float:              "Float",
	ddlparser.Float4:             "Float4",
	ddlparser.Float8:             "Float8",
	ddlparser.Double:             "Double",
	ddlparser.Real:               "Real",
	ddlparser.TinyInt:            "TinyInt",
	ddlparser.SmallInt:           "SmallInt",
	ddlparser.MediumInt:          "MediumInt",
	ddlparser.Int:                "Int",
	ddlparser.Integer:            "Integer",
	ddlparser.BigInt:             "BigInt",
	ddlparser.MiddleInt:          "MiddleInt",
	ddlparser.Int1:               "Int1",
	ddlparser.Int2:               "Int2",
	ddlparser.Int3:               "Int3",
	ddlparser.Int4:               "Int4",
	ddlparser.Int8:               "Int8",
	ddlparser.Date:               "Date",
	ddlparser.TinyBlob:           "TinyBlob",
	ddlparser.MediumBlob:         "MediumBlob",
	ddlparser.LongBlob:           "LongBlob",
	ddlparser.Bool:               "Bool",
	ddlparser.Boolean:            "Boolean",
	ddlparser.Serial:             "Serial",
	ddlparser.NVarChar:           "NVarChar",
	ddlparser.NChar:              "NChar",
	ddlparser.Char:               "Char",
	ddlparser.Character:          "Character",
	ddlparser.VarChar:            "VarChar",
	ddlparser.TinyText:           "TinyText",
	ddlparser.Text:               "Text",
	ddlparser.MediumText:         "MediumText",
	ddlparser.LongText:           "LongText",
}

func formatDataType(dt ddlparser.DataType) string {
	if dt == nil {
		return "Unknown"
	}

	name, ok := dataTypeNames[dt.Type()]
	if !ok {
		name = fmt.Sprintf("Unknown(%d)", dt.Type())
	}

	if values := dt.Value(); len(values) > 0 {
		name = fmt.Sprintf("%s[%s]", name, strings.Join(values, ", "))
	}

	if dt.Unsigned() {
		name += " Unsigned"
	}

	return name
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
