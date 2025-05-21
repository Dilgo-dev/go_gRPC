import grpc from '@grpc/grpc-js';
import protoLoader from '@grpc/proto-loader';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const PROTO_PATH = path.join(__dirname, '../../../proto/log.proto');

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true
});

const logProto = grpc.loadPackageDefinition(packageDefinition).log;

const HOST = 'localhost';
const PORT = '50051';

const client = new logProto.LogService(
  `${HOST}:${PORT}`,
  grpc.credentials.createInsecure()
);

export function createLog(message) {
  return new Promise((resolve, reject) => {
    client.createLog({ message }, (error, response) => {
      if (error) {
        reject(error);
        return;
      }
      resolve(response);
    });
  });
}

export function getLogs() {
  return new Promise((resolve, reject) => {
    client.getLogs({}, (error, response) => {
      if (error) {
        reject(error);
        return;
      }
      resolve(response.logs);
    });
  });
}

export default client;