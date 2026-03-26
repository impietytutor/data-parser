export enum LogLevel {
  DEBUG = 'debug',
  INFO = 'info',
  WARN = 'warn',
  ERROR = 'error',
}

export const DefaultLogLevel = LogLevel.INFO;

export enum DataType {
  CSV = 'csv',
  JSON = 'json',
  TSV = 'tsv',
}

export const SupportedDataTypes = [DataType.CSV, DataType.JSON, DataType.TSV];

export interface File {
  path: string;
  type: DataType;
}

export interface Config {
  logLevel?: LogLevel;
  data: File[];
}

export interface ParsedData {
  [key: string]: any[];
}