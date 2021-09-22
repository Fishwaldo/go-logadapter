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

type CopyableMap   map[string]interface{}
type CopyableSlice []interface{}

// DeepCopy will create a deep copy of this map. The depth of this
// copy is all inclusive. Both maps and slices will be considered when
// making the copy.
func (m CopyableMap) DeepCopy() map[string]interface{} {
    result := map[string]interface{}{}

    for k,v := range m {
        // Handle maps
        mapvalue,isMap := v.(map[string]interface{})
        if isMap {
            result[k] = CopyableMap(mapvalue).DeepCopy()
            continue
        }

        // Handle slices
        slicevalue,isSlice := v.([]interface{})
        if isSlice {
            result[k] = CopyableSlice(slicevalue).DeepCopy()
            continue
        }

        result[k] = v
    }

    return result
}

// DeepCopy will create a deep copy of this slice. The depth of this
// copy is all inclusive. Both maps and slices will be considered when
// making the copy.
func (s CopyableSlice) DeepCopy() []interface{} {
    result := []interface{}{}

    for _,v := range s {
        // Handle maps
        mapvalue,isMap := v.(map[string]interface{})
        if isMap {
            result = append(result, CopyableMap(mapvalue).DeepCopy())
            continue
        }

        // Handle slices
        slicevalue,isSlice := v.([]interface{})
        if isSlice {
            result = append(result, CopyableSlice(slicevalue).DeepCopy())
            continue
        }

        result = append(result, v)
    }

    return result
}