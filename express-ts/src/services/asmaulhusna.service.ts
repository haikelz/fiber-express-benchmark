import data from "../configs/asmaulhusna.json";
import { AppDataSource } from "../configs/typeorm.config";
import { AsmaulHusna } from "../entities/asmaulhusna.entity";

export class AsmaulHusnaService {
  public async getAsmaulHusna(): Promise<AsmaulHusna[]> {
    const response = await AppDataSource.getRepository(AsmaulHusna).find();
    return response;
  }

  public async insertLoopAsmaulHusna(): Promise<void> {
    const asmaulhusna = data;

    for (let i = 0; i < asmaulhusna.length; i++) {
      await AppDataSource.getRepository(AsmaulHusna)
        .create({
          arab: asmaulhusna[i].arab,
          arti: asmaulhusna[i].arti,
          latin: asmaulhusna[i].latin,
          urutan: asmaulhusna[i].urutan,
        })
        .save();
    }
  }
}
