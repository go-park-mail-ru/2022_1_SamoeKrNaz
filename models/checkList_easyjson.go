// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6d88e0beDecodePLANEXABackendModels(in *jlexer.Lexer, out *CheckLists) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(CheckLists, 0, 1)
			} else {
				*out = CheckLists{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 CheckList
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6d88e0beEncodePLANEXABackendModels(out *jwriter.Writer, in CheckLists) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v CheckLists) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6d88e0beEncodePLANEXABackendModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CheckLists) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6d88e0beEncodePLANEXABackendModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CheckLists) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6d88e0beDecodePLANEXABackendModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CheckLists) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6d88e0beDecodePLANEXABackendModels(l, v)
}
func easyjson6d88e0beDecodePLANEXABackendModels1(in *jlexer.Lexer, out *CheckList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id_cl":
			out.IdCl = uint(in.Uint())
		case "title":
			out.Title = string(in.String())
		case "id_t":
			out.IdT = uint(in.Uint())
		case "CheckListItems":
			if in.IsNull() {
				in.Skip()
				out.CheckListItems = nil
			} else {
				in.Delim('[')
				if out.CheckListItems == nil {
					if !in.IsDelim(']') {
						out.CheckListItems = make([]CheckListItem, 0, 1)
					} else {
						out.CheckListItems = []CheckListItem{}
					}
				} else {
					out.CheckListItems = (out.CheckListItems)[:0]
				}
				for !in.IsDelim(']') {
					var v4 CheckListItem
					easyjson6d88e0beDecodePLANEXABackendModels2(in, &v4)
					out.CheckListItems = append(out.CheckListItems, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6d88e0beEncodePLANEXABackendModels1(out *jwriter.Writer, in CheckList) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id_cl\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.IdCl))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"id_t\":"
		out.RawString(prefix)
		out.Uint(uint(in.IdT))
	}
	{
		const prefix string = ",\"CheckListItems\":"
		out.RawString(prefix)
		if in.CheckListItems == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.CheckListItems {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson6d88e0beEncodePLANEXABackendModels2(out, v6)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CheckList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6d88e0beEncodePLANEXABackendModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CheckList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6d88e0beEncodePLANEXABackendModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CheckList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6d88e0beDecodePLANEXABackendModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CheckList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6d88e0beDecodePLANEXABackendModels1(l, v)
}
func easyjson6d88e0beDecodePLANEXABackendModels2(in *jlexer.Lexer, out *CheckListItem) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id_clit":
			out.IdClIt = uint(in.Uint())
		case "title":
			out.Description = string(in.String())
		case "id_cl":
			out.IdCl = uint(in.Uint())
		case "id_t":
			out.IdT = uint(in.Uint())
		case "isready":
			out.IsReady = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6d88e0beEncodePLANEXABackendModels2(out *jwriter.Writer, in CheckListItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id_clit\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.IdClIt))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"id_cl\":"
		out.RawString(prefix)
		out.Uint(uint(in.IdCl))
	}
	{
		const prefix string = ",\"id_t\":"
		out.RawString(prefix)
		out.Uint(uint(in.IdT))
	}
	{
		const prefix string = ",\"isready\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsReady))
	}
	out.RawByte('}')
}
