const fs = require('fs');
const path = require('path');
const { EventEmitter } = require('events');

class DataParser extends EventEmitter {
  constructor(options = {}) {
    super();
    this.options = {
      encoding: 'utf8',
      delimiter: ',',
      ...options
    };
  }

  async parseFile(filePath) {
    try {
      const absolutePath = path.resolve(filePath);
      const fileContent = await fs.promises.readFile(absolutePath, this.options.encoding);
      const parsedData = this._parseContent(fileContent);
      this.emit('parseComplete', parsedData);
      return parsedData;
    } catch (error) {
      this.emit('error', error);
      throw error;
    }
  }

  _parseContent(content) {
    const lines = content.split('\n');
    const headers = lines[0].split(this.options.delimiter);
    const data = [];

    for (let i = 1; i < lines.length; i++) {
      if (!lines[i].trim()) continue;
      
      const values = lines[i].split(this.options.delimiter);
      const entry = {};

      for (let j = 0; j < headers.length; j++) {
        entry[headers[j].trim()] = values[j] ? values[j].trim() : null;
      }

      data.push(entry);
    }

    return data;
  }

  static validateOptions(options) {
    if (options.delimiter && typeof options.delimiter !== 'string') {
      throw new Error('Delimiter must be a string');
    }
    if (options.encoding && typeof options.encoding !== 'string') {
      throw new Error('Encoding must be a string');
    }
  }
}

module.exports = DataParser;