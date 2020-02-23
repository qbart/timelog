function _timelog_autocomplete() {
  local cur

  COMPREPLY=()
  cur=${COMP_WORDS[COMP_CWORD]}
  if [ "$COMP_CWORD" == "1" ]; then
    COMPREPLY=($(compgen -W "$(timelog autocomplete commands | fzf)" -- $cur))
  elif [ "$COMP_CWORD" == "2" ] && [ "${COMP_WORDS[1]}" == "start" ]; then
    COMPREPLY=($(compgen -W "$(timelog autocomplete qlist | fzf)" -- $cur))
  else
    COMPREPLY=()
  fi

  return 0
}

complete -F _timelog_autocomplete timelog
