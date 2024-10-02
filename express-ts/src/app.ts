import bodyParser from "body-parser";
import compression from "compression";
import cors from "cors";
import dotenv from "dotenv";
import express from "express";
import "reflect-metadata";
import { AppDataSource } from "./configs/typeorm.config";
import { AsmaulHusnaService } from "./services/asmaulhusna.service";

const app = express();
AppDataSource.initialize();
dotenv.config();

async function main() {
  app.use(bodyParser.urlencoded({ extended: false }));
  app.use(bodyParser.json());
  app.use(cors());
  app.use(compression());

  app.get("/", (_, res) => {
    res.json({
      status_code: 200,
      message: "Halo",
    });
  });

  app.get("/api/all", async (req, res) => {
    const asmaulhusnaService = new AsmaulHusnaService();
    const response = await asmaulhusnaService.getAsmaulHusna();

    if (response.length < 1) {
      asmaulhusnaService.insertLoopAsmaulHusna();
    }

    res.json({
      statusCode: 200,
      total: 99,
      data: response ? response : [],
    });
  });

  app.listen("5000", () => {
    console.log("Server berjalan di port 5000!");
  });
}

main();
