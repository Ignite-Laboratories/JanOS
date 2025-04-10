echo "Synchronizing JanOS"
echo "Ignite Laboratories"

echo "[janos]"
if [ -d ".git" ]; then
  git pull
else
  git clone "https://github.com/ignite-laboratories/JanOS"
fi

synchronize() {
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
synchronize core
synchronize fugue
synchronize glitter
synchronize host
synchronize life
synchronize spark
synchronize support
synchronize tiny