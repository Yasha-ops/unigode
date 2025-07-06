# 🧬 Unigode

**Unigode** is a Unicode-based payload generator designed to evade WAFs and XSS filters by exploiting Windows' **Best Fit Mapping** and leveraging visually similar Unicode characters.
Inspired by [Orange Tsai](https://twitter.com/orange_8361)'s research on bypassing filter-based protections using charset quirks, Unigode transforms standard payloads into stealthy, executable versions that slip past naive filtering.

---

## 🚀 Features

* 🔣 Encodes text using *Best Fit* Unicode substitutions
* 🎭 Obfuscates JavaScript payloads while preserving functionality
* 🧱 Effective against poorly configured WAFs and blacklists
* 🧾 Supports inline input (`--text`) or file-based input (`--file`)
* ⚡ Fast, minimal, binary-ready for red teaming

---

## Installation

```bash
git clone https://github.com/Yasha-ops/unigode.git
cd unigode
go build -o unigode .
```

---

## Usage

```bash
unigode --text "Text to encode"
unigode --file payload.txt
```

### Example

```bash
> ./unigode --text "<img src onerror=alert(1)>"
〈ǐｍｇ śｒｃ ơņℿｒŗℴｒ‗ǎļėřʋ⌠₁）〉
```

---

## Why does it work?

Windows' **Best Fit Mapping** automatically "normalizes" certain Unicode characters to their closest ASCII equivalents when using non-Unicode code pages (e.g., `chcp 65001` vs `1252`).
This behavior can be abused to craft payloads that:

* Appear harmless or unreadable to static filters
* Get translated back to valid JavaScript on vulnerable backends
* Bypass naive sanitization logic based on byte or char matching

This technique was notably weaponized by [Orange Tsai](https://i.blackhat.com/USA-19/Thursday/us-19-Tsai-A-New-Era-Of-SSRF-Exploiting-URL-Parser-In-Trending-Programming-Languages.pdf) to bypass SSRF and WAF protections via smart encoding tricks.
Unigode brings this idea to the client-side XSS world with instant payload generation.

---

## Pro tips

* Combine with polyglot payloads or event handler tricks for full WAF bypass
* Test across Windows/IIS setups using legacy encodings for full effect
* Use in phishing, stored XSS, or CSP-locked contexts where character shape matters

---

## Disclaimer

This tool is for educational and authorized security research purposes only.
Unauthorized use against systems without consent is strictly forbidden.

