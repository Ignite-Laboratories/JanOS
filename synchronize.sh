git pull

synchronize() {
  if [ -d "$1/.git" ]; then
    (
      cd "$1" || exit
      git pull
    )
  else
    git clone "https://github.com/ignite-laboratories/$1"
  fi
}

# Call the function for multiple repositories
synchronize core
synchronize fugue
synchronize glitter
synchronize host
synchronize life
synchronize spark
synchronize support
synchronize tiny