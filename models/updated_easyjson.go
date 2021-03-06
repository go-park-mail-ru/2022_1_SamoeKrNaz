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

func easyjson191f1b5fDecodePLANEXABackendModels(in *jlexer.Lexer, out *Updated) {
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
		case "updated":
			out.UpdatedInfo = bool(in.Bool())
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
func easyjson191f1b5fEncodePLANEXABackendModels(out *jwriter.Writer, in Updated) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"updated\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.UpdatedInfo))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Updated) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson191f1b5fEncodePLANEXABackendModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Updated) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson191f1b5fEncodePLANEXABackendModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Updated) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson191f1b5fDecodePLANEXABackendModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Updated) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson191f1b5fDecodePLANEXABackendModels(l, v)
}
func easyjson191f1b5fDecodePLANEXABackendModelsIs(in *jlexer.Lexer, out *Is_okayIn) {
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
		case "Is_okay":
			out.Is_okayInfo = bool(in.Bool())
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
func easyjson191f1b5fEncodePLANEXABackendModelsIs(out *jwriter.Writer, in Is_okayIn) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Is_okay\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Is_okayInfo))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Is_okayIn) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson191f1b5fEncodePLANEXABackendModelsIs(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Is_okayIn) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson191f1b5fEncodePLANEXABackendModelsIs(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Is_okayIn) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson191f1b5fDecodePLANEXABackendModelsIs(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Is_okayIn) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson191f1b5fDecodePLANEXABackendModelsIs(l, v)
}
func easyjson191f1b5fDecodePLANEXABackendModels1(in *jlexer.Lexer, out *ImgBoard) {
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
		case "img_board":
			out.ImgPath = string(in.String())
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
func easyjson191f1b5fEncodePLANEXABackendModels1(out *jwriter.Writer, in ImgBoard) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"img_board\":"
		out.RawString(prefix[1:])
		out.String(string(in.ImgPath))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ImgBoard) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson191f1b5fEncodePLANEXABackendModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ImgBoard) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson191f1b5fEncodePLANEXABackendModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ImgBoard) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson191f1b5fDecodePLANEXABackendModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ImgBoard) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson191f1b5fDecodePLANEXABackendModels1(l, v)
}
func easyjson191f1b5fDecodePLANEXABackendModels2(in *jlexer.Lexer, out *Deleted) {
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
		case "deleted":
			out.DeletedInfo = bool(in.Bool())
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
func easyjson191f1b5fEncodePLANEXABackendModels2(out *jwriter.Writer, in Deleted) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"deleted\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.DeletedInfo))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Deleted) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson191f1b5fEncodePLANEXABackendModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Deleted) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson191f1b5fEncodePLANEXABackendModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Deleted) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson191f1b5fDecodePLANEXABackendModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Deleted) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson191f1b5fDecodePLANEXABackendModels2(l, v)
}
func easyjson191f1b5fDecodePLANEXABackendModels3(in *jlexer.Lexer, out *Avatar) {
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
		case "avatar_path":
			out.AvatarPath = string(in.String())
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
func easyjson191f1b5fEncodePLANEXABackendModels3(out *jwriter.Writer, in Avatar) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"avatar_path\":"
		out.RawString(prefix[1:])
		out.String(string(in.AvatarPath))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Avatar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson191f1b5fEncodePLANEXABackendModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Avatar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson191f1b5fEncodePLANEXABackendModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Avatar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson191f1b5fDecodePLANEXABackendModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Avatar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson191f1b5fDecodePLANEXABackendModels3(l, v)
}
func easyjson191f1b5fDecodePLANEXABackendModels4(in *jlexer.Lexer, out *Appended) {
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
		case "appended":
			out.AppendedInfo = bool(in.Bool())
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
func easyjson191f1b5fEncodePLANEXABackendModels4(out *jwriter.Writer, in Appended) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"appended\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.AppendedInfo))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Appended) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson191f1b5fEncodePLANEXABackendModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Appended) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson191f1b5fEncodePLANEXABackendModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Appended) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson191f1b5fDecodePLANEXABackendModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Appended) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson191f1b5fDecodePLANEXABackendModels4(l, v)
}
