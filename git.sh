#!/bin/bash

username='Justin England'
email=justengland@gmail.com

git config --local user.name "$username"
git config --local user.email "$email"

echo "Git credentials updated for current session."

