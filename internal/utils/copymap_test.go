/* 
MIT License

Copyright (c) 2021 Justin Hammond

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package utils

import (
    "testing"

    "github.com/stretchr/testify/require"
)

func TestCopyMap(t *testing.T) {
    m1 := map[string]interface{}{
        "a": "bbb",
        "b": map[string]interface{}{
            "c": 123,
        },
        "c": []interface{} {
            "d", "e", map[string]interface{} {
                "f": "g",
            },
        },
    }

    m2 := CopyableMap(m1).DeepCopy()

    m1["a"] = "zzz"
    delete(m1, "b")
    m1["c"].([]interface{})[1] = "x"
    m1["c"].([]interface{})[2].(map[string]interface{})["f"] = "h"

    require.Equal(t, map[string]interface{}{
        "a": "zzz", 
        "c": []interface{} {
            "d", "x", map[string]interface{} {
                "f": "h",
            },
        },
    }, m1)
    require.Equal(t, map[string]interface{}{
        "a": "bbb",
        "b": map[string]interface{}{
            "c": 123,
        },
        "c": []interface{} {
            "d", "e", map[string]interface{} {
                "f": "g",
            },
        },
    }, m2)
}