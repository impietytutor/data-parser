// parser.js

import { join } from 'path';
import { parse } from 'csv-parse';

class Parser {
  constructor(fileContents, options = {}) {
    this.fileContents = fileContents;
    this.options = options;
    this.rows = [];
  }

  setOptions(options) {
    this.options = { ...this.options, ...options };
  }

  async parse() {
    const csvStream = parse(this.fileContents, this.options);
    const reader = csvStream.read();
    this.rows = [];

    for await (const row of reader) {
      this.rows.push(row);
    }

    return this.rows;
  }
}

export default Parser;