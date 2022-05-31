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

func easyjson9806e1DecodePLANEXABackendModels(in *jlexer.Lexer, out *Notifications) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Notifications, 0, 0)
			} else {
				*out = Notifications{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Notification
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
func easyjson9806e1EncodePLANEXABackendModels(out *jwriter.Writer, in Notifications) {
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
func (v Notifications) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9806e1EncodePLANEXABackendModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Notifications) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9806e1EncodePLANEXABackendModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Notifications) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9806e1DecodePLANEXABackendModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Notifications) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9806e1DecodePLANEXABackendModels(l, v)
}
func easyjson9806e1DecodePLANEXABackendModels1(in *jlexer.Lexer, out *Notification) {
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
		case "idu":
			out.IdU = uint(in.Uint())
		case "notification_type":
			out.NotificationType = string(in.String())
		case "date":
			out.Date = string(in.String())
		case "DateToOrder":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateToOrder).UnmarshalJSON(data))
			}
		case "is_read":
			out.IsRead = bool(in.Bool())
		case "idb":
			out.IdB = uint(in.Uint())
		case "idt":
			out.IdT = uint(in.Uint())
		case "id_wh":
			out.IdWh = uint(in.Uint())
		case "board":
			(out.Board).UnmarshalEasyJSON(in)
		case "task":
			(out.Task).UnmarshalEasyJSON(in)
		case "user_who":
			(out.UserWho).UnmarshalEasyJSON(in)
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
func easyjson9806e1EncodePLANEXABackendModels1(out *jwriter.Writer, in Notification) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"idu\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.IdU))
	}
	{
		const prefix string = ",\"notification_type\":"
		out.RawString(prefix)
		out.String(string(in.NotificationType))
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.String(string(in.Date))
	}
	{
		const prefix string = ",\"DateToOrder\":"
		out.RawString(prefix)
		out.Raw((in.DateToOrder).MarshalJSON())
	}
	{
		const prefix string = ",\"is_read\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsRead))
	}
	{
		const prefix string = ",\"idb\":"
		out.RawString(prefix)
		out.Uint(uint(in.IdB))
	}
	{
		const prefix string = ",\"idt\":"
		out.RawString(prefix)
		out.Uint(uint(in.IdT))
	}
	{
		const prefix string = ",\"id_wh\":"
		out.RawString(prefix)
		out.Uint(uint(in.IdWh))
	}
	{
		const prefix string = ",\"board\":"
		out.RawString(prefix)
		(in.Board).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"task\":"
		out.RawString(prefix)
		(in.Task).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"user_who\":"
		out.RawString(prefix)
		(in.UserWho).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Notification) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9806e1EncodePLANEXABackendModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Notification) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9806e1EncodePLANEXABackendModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Notification) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9806e1DecodePLANEXABackendModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Notification) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9806e1DecodePLANEXABackendModels1(l, v)
}
