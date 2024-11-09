/*
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
 * STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED
 * OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package main

// ANSI color codes: These constants define the escape sequences for text formatting
// and color customization in terminal-based applications (CLI). ANSI escape codes
// are used to change text styles, text colors, and background colors in supported terminals.

const (
	// Text Styling Constants
	_NC        = "\033[0m" // Reset all styles to default (e.g., resets colors, bold, underlines, etc.)
	_BOLD      = "\033[1m" // Apply bold text style
	_UNDERLINE = "\033[4m" // Apply underline to the text
	_REVERSE   = "\033[7m" // Reverse the colors (foreground becomes background and vice versa)
	_STRIKE    = "\033[9m" // Apply strikethrough to the text

	// Text Colors (foreground colors)
	_BLACK   = "\033[30m" // Set text color to black
	_RED     = "\033[31m" // Set text color to red
	_GREEN   = "\033[32m" // Set text color to green
	_YELLOW  = "\033[33m" // Set text color to yellow
	_BLUE    = "\033[34m" // Set text color to blue
	_MAGENTA = "\033[35m" // Set text color to magenta (purple)
	_CYAN    = "\033[36m" // Set text color to cyan (light blue)
	_WHITE   = "\033[37m" // Set text color to white

	// Background Colors
	_BLACKBG   = "\033[40m" // Set background color to black
	_REDBG     = "\033[41m" // Set background color to red
	_GREENBG   = "\033[42m" // Set background color to green
	_YELLOWBG  = "\033[43m" // Set background color to yellow
	_BLUEBG    = "\033[44m" // Set background color to blue
	_MAGENTABG = "\033[45m" // Set background color to magenta (purple)
	_CYANBG    = "\033[46m" // Set background color to cyan (light blue)
	_WHITEBG   = "\033[47m" // Set background color to white
)
