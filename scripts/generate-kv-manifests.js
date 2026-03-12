#!/usr/bin/env node
// generate-kv-manifests.js
// Converts brain36z markdown assets into JSON manifests for Workers KV seeding.
//
// Usage: node scripts/generate-kv-manifests.js
//
// Reads from:  ../brain36z/{playbooks,templates,workflows,agents}/*.md
// Writes to:   workers/src/assets/{playbooks,templates,workflows,agents}.json

const fs = require("fs");
const path = require("path");

const ASSET_TYPES = ["playbooks", "templates", "workflows", "agents"];
const BRAIN_DIR = path.resolve(__dirname, "../../brain36z");
const OUTPUT_DIR = path.resolve(__dirname, "../workers/src/assets");

function parseMarkdownAsset(filePath) {
  const content = fs.readFileSync(filePath, "utf-8");
  const name = path.basename(filePath, ".md");

  // Extract frontmatter-style metadata from first lines
  const lines = content.split("\n");
  let title = name;
  let description = "";

  for (const line of lines) {
    if (line.startsWith("# ")) {
      title = line.slice(2).trim();
      break;
    }
  }

  // Use first paragraph as description
  const paragraphs = content.split("\n\n");
  if (paragraphs.length > 1) {
    description = paragraphs[1].replace(/\n/g, " ").trim().slice(0, 200);
  }

  return {
    name,
    title,
    description,
    content,
    file: path.basename(filePath),
    updated: fs.statSync(filePath).mtime.toISOString(),
  };
}

function processAssetType(type_) {
  const sourceDir = path.join(BRAIN_DIR, type_);

  if (!fs.existsSync(sourceDir)) {
    console.log(`  ⊘ ${type_}: source directory not found, skipping`);
    return [];
  }

  const files = fs
    .readdirSync(sourceDir)
    .filter((f) => f.endsWith(".md"))
    .sort();

  const assets = files.map((f) => parseMarkdownAsset(path.join(sourceDir, f)));

  console.log(`  ✓ ${type_}: ${assets.length} assets`);
  return assets;
}

// Main
console.log("Generating KV manifests...\n");

fs.mkdirSync(OUTPUT_DIR, { recursive: true });

for (const type_ of ASSET_TYPES) {
  const assets = processAssetType(type_);
  const outputPath = path.join(OUTPUT_DIR, `${type_}.json`);
  fs.writeFileSync(outputPath, JSON.stringify(assets, null, 2));
}

console.log(`\nManifests written to ${OUTPUT_DIR}/`);
console.log("Seed to KV with: wrangler kv:bulk put --binding KV <file>");
