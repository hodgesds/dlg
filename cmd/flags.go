// Copyright © 2020 Daniel Hodges <hodges.daniel.scott@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	count  int
	debug  bool
	dur    time.Duration
	name   string
	repeat int
	tags   = []string{}
)

func planFlags() *pflag.FlagSet {
	planFlags := pflag.NewFlagSet("", pflag.ExitOnError)
	planFlags.BoolVar(&debug, "debug", false, "debug")
	planFlags.StringVar(&name, "name", "cli", "execution name")
	planFlags.IntVarP(&repeat, "repeat", "r", 0, "number of times to repeat")
	planFlags.IntVarP(&count, "count", "c", 100, "number of times to execute")
	planFlags.DurationVarP(&dur, "duration", "d", 0, "execution duration")
	planFlags.StringSliceVar(&tags, "tags", nil, "metrics tags")
	return planFlags
}

func flagsFromStruct(s interface{}) (*pflag.FlagSet, error) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr && v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a pointer or struct got: %t", s)
	}
	v = reflect.Indirect(v)

	nFields := v.NumField()
	flagSet := pflag.NewFlagSet(strings.ToLower(v.Type().Name()), pflag.ContinueOnError)

	for i := 0; i < nFields; i++ {
		f := v.Field(i)
		ind := reflect.Indirect(f)
		if !ind.IsValid() {
			continue
		}
		baseKind := ind.Kind()
		name := ind.Type().Name()
		isPointer := f.Kind() == reflect.Ptr
		switch baseKind {
		case reflect.Bool:
			if isPointer {
				flagSet.Bool(name, ind.Interface().(bool), "")
				continue
			}
			flagSet.Bool(name, false, "")
		case reflect.Int:
			if isPointer {
				flagSet.Int(name, ind.Interface().(int), "")
				continue
			}
			flagSet.Int(name, 0, "")
		case reflect.Int8:
			// pflag doesn't support int8 or int16
			fallthrough
		case reflect.Int16:
			if isPointer {
				flagSet.Int16(name, ind.Interface().(int16), "")
				continue
			}
			flagSet.Int16(name, 0, "")
		case reflect.Int32:
			if isPointer {
				flagSet.Int32(name, ind.Interface().(int32), "")
				continue
			}
			flagSet.Int32(name, 0, "")
		case reflect.Int64:
			if isPointer {
				flagSet.Int64(name, ind.Interface().(int64), "")
				continue
			}
			flagSet.Int64(name, 0, "")
		case reflect.Uint:
			if isPointer {
				flagSet.Uint(name, ind.Interface().(uint), "")
				continue
			}
			flagSet.Uint(name, 0, "")
		case reflect.Uint8:
			fallthrough
		case reflect.Uint16:
			if isPointer {
				flagSet.Uint16(name, ind.Interface().(uint16), "")
				continue
			}
			flagSet.Uint16(name, 0, "")
		case reflect.Uint32:
			if isPointer {
				flagSet.Uint32(name, ind.Interface().(uint32), "")
				continue
			}
			flagSet.Uint32(name, 0, "")
		case reflect.Uint64:
			if isPointer {
				flagSet.Uint64(name, ind.Interface().(uint64), "")
				continue
			}
			flagSet.Uint64(name, 0, "")
		case reflect.Float32:
			if isPointer {
				flagSet.Float32(name, ind.Interface().(float32), "")
				continue
			}
			flagSet.Uint32(name, 0, "")
		case reflect.Float64:
			if isPointer {
				flagSet.Float64(name, ind.Interface().(float64), "")
				continue
			}
			flagSet.Float64(name, 0, "")
		case reflect.String:
			if isPointer {
				flagSet.String(name, ind.Interface().(string), "")
				continue
			}
			flagSet.String(strings.ToLower(name), "", "name")
		case reflect.Array:
			continue
		case reflect.Slice:
			continue
		case reflect.Struct:
			subFlags, err := flagsFromStruct(f.Interface())
			if err != nil {
				return nil, err
			}
			flagSet.AddFlagSet(subFlags)
		}
		break
	}

	return flagSet, nil
}

func commandsFromStruct(s interface{}) ([]*cobra.Command, error) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Ptr && v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a pointer or struct got: %t", s)
	}
	v = reflect.Indirect(v)

	commands := []*cobra.Command{}

	nFields := v.NumField()
	for i := 0; i < nFields; i++ {
		f := v.Field(i)
		ind := reflect.Indirect(f)
		baseKind := ind.Kind()
		if !ind.IsValid() {
			newF := reflect.New(f.Type().Elem())
			f.Set(newF)
			ind = reflect.Indirect(reflect.Indirect(f))
			baseKind = ind.Kind()
		}
		//if baseKind != reflect.Struct && ind.CanAddr() {
		//	ind = f.Elem()
		//}
		//isPointer := f.Kind() == reflect.Ptr
		name := reflect.Indirect(ind).Type().Name()
		fmt.Printf("%+v %+v %+v\n", f, baseKind, name)

		switch baseKind {
		case reflect.Struct:
			flags, err := flagsFromStruct(f.Interface())
			if err != nil {
				return nil, err
			}
			cmd := &cobra.Command{
				Use:   strings.ToLower(name),
				Short: strings.ToLower(name),
				Long:  ``,
			}
			cmd.PersistentFlags().AddFlagSet(flags)
			commands = append(commands, cmd)
		}
	}

	return commands, nil
}
