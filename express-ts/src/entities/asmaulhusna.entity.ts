import { BaseEntity, Column, Entity, PrimaryColumn } from "typeorm";

@Entity()
export class AsmaulHusna extends BaseEntity {
  @PrimaryColumn("int")
  urutan: number;

  @Column("varchar", { length: 255 })
  latin: string;

  @Column("varchar", { length: 255 })
  arab: string;

  @Column("varchar", { length: 255 })
  arti: string;
}
