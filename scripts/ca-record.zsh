export CA_BIN="/Users/williambehn/code/projects/cmd-alytics/bin/ca"
export CA_DB="/Users/williambehn/code/projects/cmd-alytics/database/history.db"

alias ca="$CA_BIN"

autoload -Uz add-zsh-hook

#NB! Legg til ting du ikke vil skal bli logget her
typeset -a _CLIS_IGNORE
_CLIS_IGNORE+=(
  "* token *" "*apikey*" "*api_key*" "*password*" "*passwd*"
  "*secret*" "*auth*" "*--password*" "*-p*"
  "ssh *" "gpg *" "pass *"
)

_clis_preexec() {
  local line="$1"

  # skipper tomme linjer 
  [[ -z "$line" ]] && return

  # skipper space
  [[ "$line" == \ * ]] && return

  local -a words
  words=(${(z)line})
  local first=$words[1]

  # skipper egne kommandoer
  [[ $first == $CA_BIN || $first == "ca" ]] && return

  # skipper sudo kommandoer
  [[ $first == sudo ]] && return

  # ignorer ting som ikke skal bli logget
  local pat
  for pat in "${_CLIS_IGNORE[@]}"; do
    [[ "$line" == $pat ]] && return
  done

  local dir="$PWD"
  local -i ts=$EPOCHSECONDS
  local shell="zsh"

  # json laget med jq pipes til record cmd. Non blocking (&) og gir ingen return verdi (!)
  if command -v jq >/dev/null 2>&1; then
    { jq -n --arg cmd "$line" --arg dir "$dir" --argjson ts "$ts" \
         '{cmd:$cmd, dir:$dir, ts:$ts}' \
      | "$CA_BIN" record; } >/dev/null 2>&1 &!
  else
    # Hvis jq ikke er installert, idk 
    :
  fi
}

add-zsh-hook preexec _clis_preexec
