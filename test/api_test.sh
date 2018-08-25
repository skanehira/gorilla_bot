#!/bin/bash

# post url_verification
echo "#### url_verification"
curl -s -v -H "Content-type: application/json" http://localhost:8080/slack/gorilla -d @url_verification.json

echo ""
echo "#### member joined channel"
# post member joined channel
curl -s -v -H "Content-type: application/json" http://localhost:8080/slack/gorilla -d @member_joined_channel.json
