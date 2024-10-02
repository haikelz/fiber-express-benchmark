import dotenv from "dotenv";
import { DataSource } from "typeorm";
import { AsmaulHusna } from "../entities/asmaulhusna.entity";
import {
  DATABASE_HOST,
  DATABASE_NAME,
  DATABASE_PASSWORD,
  DATABASE_PORT,
  DATABASE_USERNAME,
} from "./constants.config";

dotenv.config();

export const AppDataSource = new DataSource({
  type: "postgres",
  host: DATABASE_HOST,
  port: Number(DATABASE_PORT),
  username: DATABASE_USERNAME,
  password: DATABASE_PASSWORD,
  database: DATABASE_NAME,
  entities: [AsmaulHusna],
  synchronize: true,
  logging: true,
  extra: {
    connectionLimit: 50000
  }
});
