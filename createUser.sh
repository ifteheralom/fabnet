#!/bin/bash

set -e

sudo rm -rf hfc-key-store
node enrollAdmin.js
node registerUser.js
node query.js
