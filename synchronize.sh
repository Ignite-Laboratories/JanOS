SCRIPT_DIR=$(dirname "$(realpath "$0")")
(
  cd "$SCRIPT_DIR" || exit

  echo "Synchronizing JanOS"
  echo "Ignite Laboratories"

  echo

  echo "[janos]"
  if [ -d ".git" ]; then
    git pull
  else
    git clone "https://github.com/ignite-laboratories/janos"
  fi

  synchronize() {
    echo
    if [ -d "$1/.git" ]; then
      (
        cd "$1" || exit
        echo "[$1]"
        git pull
      )
    else
      echo "[$1]"
      git clone "https://github.com/ignite-laboratories/$1"
    fi
  }

  # Call the function for multiple repositories
  synchronize arwen
  synchronize core
  synchronize fugue
  synchronize glitter
  synchronize host
  synchronize hydra
  synchronize life
  synchronize spark
  synchronize support
  synchronize tiny
)