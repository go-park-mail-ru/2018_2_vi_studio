// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package game

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

func easyjson66c1e240DecodeGameCore(in *jlexer.Lexer, out *msgWrongTry) {
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
func easyjson66c1e240EncodeGameCore(out *jwriter.Writer, in msgWrongTry) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v msgWrongTry) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson66c1e240EncodeGameCore(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v msgWrongTry) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson66c1e240EncodeGameCore(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *msgWrongTry) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson66c1e240DecodeGameCore(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *msgWrongTry) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson66c1e240DecodeGameCore(l, v)
}
func easyjson66c1e240DecodeGameCore1(in *jlexer.Lexer, out *msgQueuePosition) {
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
		case "position":
			out.Position = int(in.Int())
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
func easyjson66c1e240EncodeGameCore1(out *jwriter.Writer, in msgQueuePosition) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"position\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Position))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v msgQueuePosition) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson66c1e240EncodeGameCore1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v msgQueuePosition) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson66c1e240EncodeGameCore1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *msgQueuePosition) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson66c1e240DecodeGameCore1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *msgQueuePosition) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson66c1e240DecodeGameCore1(l, v)
}
func easyjson66c1e240DecodeGameCore2(in *jlexer.Lexer, out *msgPlayer) {
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
		case "id":
			out.Id = int(in.Int())
		case "nickname":
			out.Nickname = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
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
func easyjson66c1e240EncodeGameCore2(out *jwriter.Writer, in msgPlayer) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Id))
	}
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
		const prefix string = ",\"avatar\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Avatar))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v msgPlayer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson66c1e240EncodeGameCore2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v msgPlayer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson66c1e240EncodeGameCore2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *msgPlayer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson66c1e240DecodeGameCore2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *msgPlayer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson66c1e240DecodeGameCore2(l, v)
}
func easyjson66c1e240DecodeGameCore3(in *jlexer.Lexer, out *msgNextTry) {
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
		case "lastTry":
			(out.LastTry).UnmarshalEasyJSON(in)
		case "currentTry":
			(out.CurrentTry).UnmarshalEasyJSON(in)
		case "gameOver":
			(out.GameOver).UnmarshalEasyJSON(in)
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
func easyjson66c1e240EncodeGameCore3(out *jwriter.Writer, in msgNextTry) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"lastTry\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.LastTry).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"currentTry\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.CurrentTry).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"gameOver\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.GameOver).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v msgNextTry) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson66c1e240EncodeGameCore3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v msgNextTry) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson66c1e240EncodeGameCore3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *msgNextTry) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson66c1e240DecodeGameCore3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *msgNextTry) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson66c1e240DecodeGameCore3(l, v)
}
func easyjson66c1e240DecodeGameCore4(in *jlexer.Lexer, out *msgGameStart) {
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
		case "userId":
			out.UserId = int(in.Int())
		case "players":
			if in.IsNull() {
				in.Skip()
				out.Players = nil
			} else {
				in.Delim('[')
				if out.Players == nil {
					if !in.IsDelim(']') {
						out.Players = make([]msgPlayer, 0, 1)
					} else {
						out.Players = []msgPlayer{}
					}
				} else {
					out.Players = (out.Players)[:0]
				}
				for !in.IsDelim(']') {
					var v1 msgPlayer
					(v1).UnmarshalEasyJSON(in)
					out.Players = append(out.Players, v1)
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
func easyjson66c1e240EncodeGameCore4(out *jwriter.Writer, in msgGameStart) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.UserId))
	}
	{
		const prefix string = ",\"players\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Players == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Players {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v msgGameStart) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson66c1e240EncodeGameCore4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v msgGameStart) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson66c1e240EncodeGameCore4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *msgGameStart) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson66c1e240DecodeGameCore4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *msgGameStart) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson66c1e240DecodeGameCore4(l, v)
}
func easyjson66c1e240DecodeGameCore5(in *jlexer.Lexer, out *msgGameOver) {
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
func easyjson66c1e240EncodeGameCore5(out *jwriter.Writer, in msgGameOver) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v msgGameOver) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson66c1e240EncodeGameCore5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v msgGameOver) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson66c1e240EncodeGameCore5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *msgGameOver) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson66c1e240DecodeGameCore5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *msgGameOver) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson66c1e240DecodeGameCore5(l, v)
}
func easyjson66c1e240DecodeGameCore6(in *jlexer.Lexer, out *msgDoneTry) {
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
		case "row":
			out.Row = int(in.Int())
		case "col":
			out.Col = int(in.Int())
		case "rotationCount":
			out.Rotation = int(in.Int())
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
func easyjson66c1e240EncodeGameCore6(out *jwriter.Writer, in msgDoneTry) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"row\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Row))
	}
	{
		const prefix string = ",\"col\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Col))
	}
	{
		const prefix string = ",\"rotationCount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Rotation))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v msgDoneTry) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson66c1e240EncodeGameCore6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v msgDoneTry) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson66c1e240EncodeGameCore6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *msgDoneTry) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson66c1e240DecodeGameCore6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *msgDoneTry) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson66c1e240DecodeGameCore6(l, v)
}
func easyjson66c1e240DecodeGameCore7(in *jlexer.Lexer, out *msgCurretTry) {
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
		case "playerId":
			out.PlayerId = int(in.Int())
		case "tileType":
			out.TileType = int(in.Int())
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
func easyjson66c1e240EncodeGameCore7(out *jwriter.Writer, in msgCurretTry) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"playerId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.PlayerId))
	}
	{
		const prefix string = ",\"tileType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.TileType))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v msgCurretTry) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson66c1e240EncodeGameCore7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v msgCurretTry) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson66c1e240EncodeGameCore7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *msgCurretTry) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson66c1e240DecodeGameCore7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *msgCurretTry) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson66c1e240DecodeGameCore7(l, v)
}
func easyjson66c1e240DecodeGameCore8(in *jlexer.Lexer, out *Message) {
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
		case "event":
			out.Event = string(in.String())
		case "data":
			if m, ok := out.Data.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Data.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Data = in.Interface()
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
func easyjson66c1e240EncodeGameCore8(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"event\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Event))
	}
	{
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if m, ok := in.Data.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Data.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Data))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson66c1e240EncodeGameCore8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson66c1e240EncodeGameCore8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson66c1e240DecodeGameCore8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson66c1e240DecodeGameCore8(l, v)
}
