// types.ts

export interface DataParserConfig {
  inputFilePath: string;
  outputFilePath: string;
  delimiter: string;
}

export interface DataRow {
  id: number;
  name: string;
  values: number[];
}

export interface ParsedData {
  data: DataRow[];
  errors: string[];
}

export enum ParserStatus {
  PENDING = 'pending',
  IN_PROGRESS = 'in_progress',
  COMPLETED = 'completed',
  FAILED = 'failed',
}

export interface ParserResult {
  status: ParserStatus;
  data: ParsedData | null;
  errorMessage: string | null;
}