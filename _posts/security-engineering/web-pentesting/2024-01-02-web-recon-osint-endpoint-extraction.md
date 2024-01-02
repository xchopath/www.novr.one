---
layout: post
title: "OSINT for Web Endpoint Extraction"
date: 2024-01-02 09:00:00 +0700
categories: "Web-Pentesting"
---

## 1. Wayback Machine (web.archive.org)

```bash
echo "target.com" | xargs -I {} curl -s "https://web.archive.org/cdx/search/cdx?url={}/*&output=text&fl=original&collapse=urlkey"
```

## 2. URLScan.io
```bash
echo "target.com" | xargs -I {} curl -s "https://urlscan.io/api/v1/search/?q=domain:{}&size=100" | jq -r '.results[].page.url' | sort -V | uniq
```

## 3. AlienVault Open Threat Exchange (otx.alienvault.com)

```bash
echo "target.com" | xargs -I {} curl -s "https://otx.alienvault.com/api/v1/indicators/domain/{}/url_list?limit=500&page=1" | jq -r '.url_list[].url'
```

## 4. CommonCrawl.org
```bash
target="target.com"
for commoncrawl in $(curl -s "http://index.commoncrawl.org/collinfo.json" | jq -r '.[]."cdx-api"')
do
  curl -s "${commoncrawl}?url=*.${target}&output=json" | jq -r .url
done
```