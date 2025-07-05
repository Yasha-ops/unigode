# 🧬 Unigode

**Unigode** est un outil de génération de payloads Unicode obfusqués, conçu pour contourner les filtres WAF/XSS à l’aide de substitutions intelligentes de caractères.  
Il encode intelligemment votre input pour le rendre toujours exécutable… mais beaucoup plus difficile à détecter.

> 🔥 *"Il pénètre les filtres comme un charme."*

---

## 🚀 Fonctionnalités

- Encode automatiquement du texte avec des caractères Unicode "best fit"
- Bypass de filtres basiques et WAFs
- Support du texte en ligne de commande ou depuis un fichier
- Rapide, minimaliste, efficace

---

## 📦 Installation

Clonez simplement le repo et compilez si besoin :

```bash
git clone https://github.com/yourusername/unigode.git
cd unigode
go build -o unigode .


## 🛠️Usage
```bash
unigode --text "Text to encode"
unigode --file payload.txt
```

### Exemple
```bash
> unigode --text "<img src onerror=alert(1)>"
〈ǐｍｇ śｒｃ ơņℿｒŗℴｒ‗ǎļėřʋ⌠₁）〉
```

## 🧠 Why doest it work ?
Unigode repose sur l’obfuscation par substitution de caractères Unicode visuellement similaires.
Cela permet de contourner des protections naïves ou mal configurées, tout en gardant l’exécution JavaScript valide dans de nombreux contextes.


