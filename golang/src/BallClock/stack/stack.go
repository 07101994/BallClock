// Copyright © 2010-12 Qtrac Ltd.
// 
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Modified from "Programming in Go" (Mark Summerfield; Addison Wesley, May 2012)
// by Mark J. Miller.

package stack

import "errors"

type Stack []int

func (stack *Stack) Pop() (int, error) {
    theStack := *stack
    if len(theStack) == 0 {
        return 0, errors.New("can't Pop() an empty stack")
    }
    x := theStack[len(theStack)-1]
    *stack = theStack[:len(theStack)-1]
    return x, nil
}

func (stack *Stack) Push(x int) {
    *stack = append(*stack, x)
}

func (stack Stack) Top() (int, error) {
    if len(stack) == 0 {
        return 0, errors.New("can't Top() an empty stack")
    }
    return stack[len(stack)-1], nil
}

func (stack Stack) Cap() int {
    return cap(stack)
}

func (stack Stack) Len() int {
    return len(stack)
}

func (stack Stack) IsEmpty() bool {
    return len(stack) == 0
}
