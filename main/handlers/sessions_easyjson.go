// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package handlers

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

func easyjsonC50d1aafDecodeMainHandlers(in *jlexer.Lexer, out *SessionHandler) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
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
func easyjsonC50d1aafEncodeMainHandlers(out *jwriter.Writer, in SessionHandler) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SessionHandler) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC50d1aafEncodeMainHandlers(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SessionHandler) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC50d1aafEncodeMainHandlers(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SessionHandler) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC50d1aafDecodeMainHandlers(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SessionHandler) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC50d1aafDecodeMainHandlers(l, v)
}
func easyjsonC50d1aafDecodeMainHandlers1(in *jlexer.Lexer, out *CreateSessionResponse) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "nickname":
			out.Nickname = string(in.String())
		case "email":
			out.Email = string(in.String())
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
func easyjsonC50d1aafEncodeMainHandlers1(out *jwriter.Writer, in CreateSessionResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nickname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateSessionResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC50d1aafEncodeMainHandlers1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateSessionResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC50d1aafEncodeMainHandlers1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateSessionResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC50d1aafDecodeMainHandlers1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateSessionResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC50d1aafDecodeMainHandlers1(l, v)
}
func easyjsonC50d1aafDecodeMainHandlers2(in *jlexer.Lexer, out *CreateSessionRequest) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "nickname":
			out.Nickname = string(in.String())
		case "password":
			out.Password = string(in.String())
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
func easyjsonC50d1aafEncodeMainHandlers2(out *jwriter.Writer, in CreateSessionRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nickname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateSessionRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC50d1aafEncodeMainHandlers2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateSessionRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC50d1aafEncodeMainHandlers2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateSessionRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC50d1aafDecodeMainHandlers2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateSessionRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC50d1aafDecodeMainHandlers2(l, v)
}
