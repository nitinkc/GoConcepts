/* Adds a "Run" button to every fenced ```go code block and executes it
   via the public Wandbox API (https://wandbox.org). */
(function () {
  var WANDBOX_API = "https://wandbox.org/api/compile.json";
  var COMPILER = "go-1.23.2"; // newest Go on wandbox.org (see /api/list.json)

  function wrapSource(code) {
    // Lesson docs contain fragments; make them self-contained when possible.
    if (!/\bpackage\s+\w+/.test(code)) {
      if (/\bfunc\s+main\s*\(/.test(code)) {
        code = "package main\n\n" + code;
      } else {
        code = "package main\n\nimport \"fmt\"\n\nfunc main() {\n" + code + "\n}";
      }
    }
    return code;
  }

  function run(codeEl, outputEl, button) {
    button.disabled = true;
    button.textContent = "Running…";
    outputEl.textContent = "";
    outputEl.hidden = false;

    fetch(WANDBOX_API, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        compiler: COMPILER,
        code: wrapSource(codeEl.innerText),
      }),
    })
      .then(function (r) { return r.json(); })
      .then(function (res) {
        var out = "";
        if (res.compiler_error) out += "[build error]\n" + res.compiler_error;
        if (res.program_output) out += res.program_output;
        if (res.program_error) out += "[runtime]\n" + res.program_error;
        if (!out) out = res.compiler_message || "(no output)";
        outputEl.textContent = out;
        button.textContent = "Run again";
      })
      .catch(function (err) {
        outputEl.textContent = "Request failed: " + err.message;
        button.textContent = "Run again";
      })
      .finally(function () {
        button.disabled = false;
      });
  }

  function enhance(codeEl) {
    if (codeEl.dataset.goRunner) return;
    codeEl.dataset.goRunner = "1";

    var host = codeEl.closest("div.highlight") || codeEl.closest("pre") || codeEl.parentElement;
    var originalCode = codeEl.innerText;
    codeEl.setAttribute("contenteditable", "plaintext-only");
    codeEl.setAttribute("role", "textbox");
    codeEl.setAttribute("aria-label", "Editable Go code");
    codeEl.setAttribute("aria-multiline", "true");
    codeEl.setAttribute("spellcheck", "false");
    codeEl.classList.add("go-runner-editor");

    var bar = document.createElement("div");
    bar.className = "go-runner-bar";

    var resetButton = document.createElement("button");
    resetButton.className = "go-runner-btn go-runner-reset";
    resetButton.type = "button";
    resetButton.textContent = "Reset";

    var button = document.createElement("button");
    button.className = "go-runner-btn";
    button.type = "button";
    button.textContent = "Run";

    var output = document.createElement("pre");
    output.className = "go-runner-output";
    output.hidden = true;

    bar.appendChild(resetButton);
    bar.appendChild(button);
    host.appendChild(bar);
    host.appendChild(output);

    codeEl.addEventListener("keydown", function (event) {
      if (event.key !== "Tab") return;
      event.preventDefault();
      var selection = window.getSelection();
      if (!selection.rangeCount) return;
      var range = selection.getRangeAt(0);
      range.deleteContents();
      var spaces = document.createTextNode("    ");
      range.insertNode(spaces);
      range.setStartAfter(spaces);
      range.collapse(true);
      selection.removeAllRanges();
      selection.addRange(range);
    });

    resetButton.addEventListener("click", function () {
      codeEl.textContent = originalCode;
      output.textContent = "";
      output.hidden = true;
      button.textContent = "Run";
    });

    button.addEventListener("click", function () {
      run(codeEl, output, button);
    });
  }

  function isRunnable(codeEl) {
    var block = codeEl.closest(".tabbed-block");
    var tabSet = codeEl.closest(".tabbed-set");
    if (!block || !tabSet) return false;

    var blocks = Array.prototype.slice.call(tabSet.querySelectorAll(":scope > .tabbed-content > .tabbed-block"));
    var labels = tabSet.querySelectorAll(":scope > .tabbed-labels > label");
    var label = labels[blocks.indexOf(block)];
    return label && label.textContent.trim().toLowerCase() === "runnable example";
  }

  function scan() {
    // pymdownx puts the language class on the wrapper div: <div class="language-go highlight">
    document
      .querySelectorAll("div.language-go pre code, div.language-golang pre code")
      .forEach(function (codeEl) {
        if (isRunnable(codeEl)) enhance(codeEl);
      });
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", scan);
  } else {
    scan();
  }
  // Re-scan when Material's instant navigation swaps content.
  document.addEventListener("DOMContentLoaded", function () {
    var observer = new MutationObserver(scan);
    observer.observe(document.body, { childList: true, subtree: true });
  });
})();
