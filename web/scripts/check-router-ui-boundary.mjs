import fs from 'node:fs';
import path from 'node:path';

const projectRoot = path.resolve(process.cwd());
const srcRoot = path.join(projectRoot, 'src');
const routerUIRoot = path.join(srcRoot, 'router-ui');
const exts = new Set(['.js', '.jsx', '.ts', '.tsx']);

const violations = [];

function walk(dirPath) {
  const entries = fs.readdirSync(dirPath, { withFileTypes: true });
  for (const entry of entries) {
    const fullPath = path.join(dirPath, entry.name);
    if (entry.isDirectory()) {
      walk(fullPath);
      continue;
    }
    if (!exts.has(path.extname(entry.name))) {
      continue;
    }
    if (fullPath.startsWith(routerUIRoot + path.sep)) {
      continue;
    }
    const content = fs.readFileSync(fullPath, 'utf8');
    const hasAntdImport =
      /from\s+['"]antd(?:\/[^'"]*)?['"]/.test(content) ||
      /import\s+['"]antd(?:\/[^'"]*)?['"]/.test(content);
    if (!hasAntdImport) {
      continue;
    }
    const relativePath = path.relative(projectRoot, fullPath);
    violations.push(relativePath);
  }
}

walk(srcRoot);

if (violations.length > 0) {
  console.error(
    [
      'Direct antd imports are only allowed inside src/router-ui.',
      'Move usage behind router-ui primitives/patterns before importing it in business code.',
      '',
      ...violations.map((item) => `- ${item}`),
    ].join('\n'),
  );
  process.exit(1);
}

console.log('router-ui boundary check passed.');
