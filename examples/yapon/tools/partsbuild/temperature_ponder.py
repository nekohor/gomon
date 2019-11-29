from base_ponder import BasePonder


class TemperaturePonder(BasePonder):

    def __init__(self):

        super().__init__()

        self.temp_tags = ["pet", "r1dt", "r2dt", "fet", "fdt", "ct"]

        self.signals = {}

        for tag in self.temp_tags:
            self.signals[tag] = {}

        self.signals["pet"]["2250"] = {
            "pet1": "TN\\L_SSP_PET1TEMP",
            "pet2": "TN\\L_SSP_PET2TEMP",
            "pet_trans_length": "TN\\L_SSP_PETTRANSLEN",
            "pet_strip_speed": "TN\\L_SSP_PETSTRIPSFB",
        }

        self.signals["pet"]["1580"] = {
            "pet_on": "TN\\L2_SSP_PETON",
            "pet_trans_length": "TN\\L2_SSP_PETTRANSLEN",
            "pet_strip_speed": "TN\\L2_SSP_PETSTRIPSFB",
            "pet0": "TN\\L2_SSP_PETTEMP",
            "pet1": "TN\\L2_SSP_PET1TEMP",
            "pet2": "TN\\L2_SSP_PET2TEMP",
            "pet_point_detect1": "TN\\L2_SSP_PETSPNT1",
            "pet_point_detect2": "TN\\L2_SSP_PETSPNT2",
            "pet_point_detect3": "TN\\L2_SSP_PETSPNT3",
            "pet_sample_point1": "TN\\L2_HSBS_SMP1PET",
            "pet_sample_point2": "TN\\L2_HSBS_SMP2PET",
            "pet_sample_point3": "TN\\L2_HSBS_SMP3PET",
            "pet_selection": "TN\\L2_SSP_PETSEL",
            "pet1_healthy": "TN\\L2_SSP_PET1HTY",
            "pet2_healthy": "TN\\L2_SSP_PET2HTY",
        }

        self.signals["r1dt"]["2250"] = {
            "r1dt1": "TN\\L_R1_R1DT1TEMP",
            "r1dt2": "TN\\L_R1_R1DT2TEMP",
            "r1dt_trans_length": "TN\\L_R1_R1DTTRANSLEN",
            "r1dt_strip_speed": "TN\\L_R1_R1DTSTRIPSFB",
        }

        self.signals["r1dt"]["1580"] = {
            "r1dt_lrx": "TN\\L2_R1_R1LRX",
            "r1dt_on": "TN\\L2_R1_R1DTON",
            "r1dt_trans_length": "TN\\L2_R1_R1DTTRANSLEN",
            "r1dt_strip_speed": "TN\\L2_R1_R1DTSTRIPSFB",
            "r1dt0": "TN\\L2_R1_R1DTTEMP",
            "r1dt1": "TN\\L2_R1_R1DT1TEMP",
            "r1dt2": "TN\\L2_R1_R1DT2TEMP",
            "r1dt_selection": "TN\\L2_R1_R1DTSEL",
            "r1dt1_healthy": "TN\\L2_R1_R1DT1HTY",
            "r1dt2_healthy": "TN\\L2_R1_R1DT2HTY",
        }

        self.signals["r2dt"]["2250"] = {
            "r2dt1": "TN\\L_R2_R2DT1TEMP",
            "r2dt2": "TN\\L_R2_R2DT2TEMP",
            "r2dt_trans_length": "TN\\L_R2_R2DTTRANSLEN",
            "r2dt_strip_speed": "TN\\L_R2_R2DTSTRIPSFB",
        }

        self.signals["r2dt"]["1580"] = {
            "r2dt_lrx": "TN\\L2_R2_R2LRX",
            "r2dt_on": "TN\\L2_R2_R2DTON",
            "r2dt_trans_length": "TN\\L2_R2_R2DTTRANSLEN",
            "r2dt_strip_speed": "TN\\L2_R2_R2DTSTRIPSFB",
            "r2dt0": "TN\\L2_R2_R2DTTEMP",
            "r2dt1": "TN\\L2_R2_R2DT1TEMP",
            "r2dt2": "TN\\L2_R2_R2DT2TEMP",
            "r2dt_point_detect1": "TN\\L2_R2_R2DTSPNT1",
            "r2dt_point_detect2": "TN\\L2_R2_R2DTSPNT2",
            "r2dt_point_detect3": "TN\\L2_R2_R2DTSPNT3",
            "r2dt_sample_point1": "TN\\L2_R2S_SPTR2DT1",
            "r2dt_sample_point2": "TN\\L2_R2S_SPTR2DT2",
            "r2dt_sample_point3": "TN\\L2_R2S_SPTR2DT3",
            "r2dt_selection": "TN\\L2_R2_R2DTSEL",
            "r2dt1_healthy": "TN\\L2_R2_R2DT1HTY",
            "r2dt2_healthy": "TN\\L2_R2_R2DT2HTY",
        }

        self.signals["fet"]["2250"] = {
            "fet1": "TN\\L_FME_FET1TEMP",
            "fet2": "TN\\L_FME_FET2TEMP",
            "fet_trans_length": "TN\\L_FME_TRANSLEN",
            "fet_strip_speed": "TN\\L_FME_FETSTRIPSFB",
        }

        self.signals["fet"]["1580"] = {
            "fet_on": "TN\\L2_FME_FETON",
            "fet_trans_length": "TN\\L2_FME_TRANSLEN",
            "fet_strip_speed": "TN\\L2_FME_FETSTRIPSFB",
            "fet0": "TN\\L2_FME_FETTEMP",
            "fet1": "TN\\L2_FME_FET1TEMP",
            "fet2": "TN\\L2_FME_FET2TEMP",
            "fet_selection": "TN\\L2_FME_FETSEL",
            "fet1_ready": "TN\\L2_FME_FET_RDY1",
            "fet2_ready": "TN\\L2_FME_FET_RDY2",
        }

        self.signals["fdt"]["2250"] = {
            "fdt1": "TN\\L_FA_FDT1TEMP",
            "fdt2": "TN\\L_FA_FDT2TEMP",
            "fdt_trans_length": "TN\\L_FM_FDTTRANSLEN",
        }

        self.signals["fdt"]["1580"] = {
            "fdt_metal_detect": "TN\\L2_ROT_FDT",
            "fdt_trans_length": "TN\\L2_ROT_FDTTRFLEN",
            "fdt_strip_speed": "TN\\L2_ROT_FDTSTRIPSFB",
            "fdt0": "TN\\L2_FA_FDTTEMP",
            "fdt1": "TN\\L2_FA_FDT1TEMP",
            "fdt2": "TN\\L2_FA_FDT2TEMP",
            "fdt_selection": "TN\\L2_FA_FDTSEL",
            "fdt1_healthy": "TN\\L2_FA_FDT1HTY",
            "fdt2_healthy": "TN\\L2_FA_FDT2HTY",
            "fdt_point_detect1": "TN\\L2_ROT_HESMPRCHFDT1",
            "fdt_point_detect2": "TN\\L2_ROT_HESMPRCHFDT2",
            "fdt_point_detect3": "TN\\L2_ROT_HESMPRCHFDT3",
            "head_end_dc1": "TN\\L2_ROT_HEDC1Z",
            "head_end_dc2": "TN\\L2_ROT_HEDC2Z",
        }

        self.signals["ct"]["2250"] = {
            "ct1": "TN\\L_ROT_CT1TEMP",
            "ct2": "TN\\L_ROT_CT2TEMP",
            "ct_strip_speed": "TN\\L_ROT_CTSTRIPSFB",
            "ct_trans_length": "TN\\L_ROT_CTTRANSLEN",
            "ct_tgt": "PR\\CTC_CT_CLS_TGT",
            "cls_permit": "PR\\CTC_CT_CLS_PERMIT",
            "ct_selection": "TN\\L_ROT_CTSEL",
        }

        self.signals["ct"]["1580"] = {
            "ct_metal_detect": "TN\\L2_ROT_CT",
            "ct_trans_length": "TN\\L2_ROT_CTTRANSLEN",
            "ct0": "TN\\L2_ROT_CTTMP",
            "ct1": "TN\\L2_ROT_CT1TEMP",
            "ct2": "TN\\L2_ROT_CT2TEMP",
            "ct_strip_speed": "TN\\L2_ROT_CTSTRIPSFB",
            "ct_tgt": "PR\\CTC_CT_CLS_TGT",
            "cls_permit": "PR\\CTC_CT_CLS_PERMIT",
            "ct_selection": "TN\\L2_ROT_CTSEL",
            "ct1_healthy": "TN\\L2_ROT_CT1HTY",
            "ct2_healthy": "TN\\L2_ROT_CT2HTY",
            "ct_point_detect1": "TN\\L2_ROT_SMP1RCHCT",
            "ct_point_detect2": "TN\\L2_ROT_SMP2RCHCT",
            "ct_point_detect3": "TN\\L2_ROT_SMP3RCHCT",
        }

    def build_df(self):
        for line in self.lines:
            for tag in self.temp_tags:
                for part, signal in self.signals[tag][line].items():

                    if part == tag + "1":
                        self.df = self.df.append({
                            "LINE": line,
                            "PART": tag,
                            "DCAFILE": self.get_dca_file_name(tag),
                            "SIGNAL": signal
                        }, ignore_index=True)

                    self.df = self.df.append({
                        "LINE": line,
                        "PART": part,
                        "DCAFILE": self.get_dca_file_name(tag),
                        "SIGNAL": signal
                    }, ignore_index=True)
