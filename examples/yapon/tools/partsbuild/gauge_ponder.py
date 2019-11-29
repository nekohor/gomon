from base_ponder import BasePonder


class GaugePonder(BasePonder):

    def __init__(self):

        super().__init__()

        self.signals = {}

        self.signals["2250"] = {
            "thick_clg": "TN\\L_AG2_F7XTHKCDEVCLG",
            "thick_mfg": "TN\\L_AG2_F7XTHKCDEVMFG",
            "width_flt": "TN\\L_FA_F7XWIDDEVFLT",
            "width_mfg": "TN\\L_FA_F7XWIDDEVMFG",

            "wedge25": "L2\\FMX_MFT_PROFB_WEDGE1",
            "wedge40": "L2\\FMX_MFT_PROFB_WEDGE2",
            "wedge70": "L2\\FMX_MFT_PROFB_WEDGE3",
            "wedge100": "L2\\FMX_MFT_PROFB_WEDGE4",

            "crown25": "L2\\FMX_MFT_PROFB_CROWN1",
            "crown40": "L2\\FMX_MFT_PROFB_CROWN2",
            "crown70": "L2\\FMX_MFT_PROFB_CROWN3",
            "crown100": "L2\\FMX_MFT_PROFB_CROWN4",

            "flt_ro1": "TN\\L_CF_F7XFLTRO1",
            "flt_ro2": "TN\\L_CF_F7XFLTRO2",
            "flt_ro3": "TN\\L_CF_F7XFLTRO3",
            "flt_ro4": "TN\\L_CF_F7XFLTRO4",
            "flt_ro5": "TN\\L_CF_F7XFLTRO5",

            "sym_flt_del": "TN\\L_CF_F7XFLT",

            "r2_cent_off": "TN\\L_R2_R2XCENTOFS",
            "fm_cent_off": "TN\\L_FA_F7XCENTOFSFLT",
        }

        self.signals["1580"] = {
            "thick_clg": "TN\\L2_AGC_F7XTHKCDEVCLG",
            "thick_mfg": "TN\\L2_AGC_F7XTHKCDEVAGC",
            "width_flt": "TN\\L2_FA_FLTWDEV",
            "width_mfg": "TN\\L2_CF_F7XWIDDEVMFG",

            "wedge25": "L2\\MFG_PROFB_WEDGE1",
            "wedge40": "L2\\MFG_PROFB_WEDGE2",
            "wedge70": "L2\\MFG_PROFB_WEDGE3",
            "wedge100": "L2\\MFG_PROFB_WEDGE4",

            "crown25": "L2\\MFG_PROFB_CROWN1",
            "crown40": "L2\\MFG_PROFB_CROWN2",
            "crown70": "L2\\MFG_PROFB_CROWN3",
            "crown100": "L2\\MFG_PROFB_CROWN4",

            "wedge25_fit": "PR\\FM_DEL_WEDGE25",
            "wedge40_fit": "PR\\FM_DEL_WEDGE40",
            "wedge70_fit": "PR\\FM_DEL_WEDGE70",
            "wedge100_fit": "PR\\FM_DEL_WEDGE100",

            "crown25_fit": "PR\\FM_DEL_CROWN25",
            "crown40_fit": "PR\\FM_DEL_CROWN40",
            "crown70_fit": "PR\\FM_DEL_CROWN70",
            "crown100_fit": "PR\\FM_DEL_CROWN100",

            "flt_ro1": "TN\\L2_CF_F7XFLTRO1",
            "flt_ro2": "TN\\L2_CF_F7XFLTRO2",
            "flt_ro3": "TN\\L2_CF_F7XFLTRO3",
            "flt_ro4": "TN\\L2_CF_F7XFLTRO4",
            "flt_ro5": "TN\\L2_CF_F7XFLTRO5",

            "sym_flt_del": "TN\\L2_CF_F7XFLTSYM",
            "asym_flt_del": "TN\\L2_CF_F7XFLTASYM",

            "r2_cent_off": "TN\\L2_R2_R2XCENTOFS",
            "fm_cent_off": "TN\\L2_CF_F7XPRFHKDEV",
        }

    def get_gauge_dca_file_name(self, line, part):

        if part[:5] == "thick" or part[:5] == "width":
            return self.get_dca_file_name(part.split("_")[-1])

        elif part[:5] == "wedge" or part[:5] == "crown":
            return self.get_dca_file_name("mfg")

        elif part[:3] == "flt":
            return self.get_dca_file_name("flt")

        elif part[-3:] == "del":
            return self.get_dca_file_name("flt")

        elif part == "fm_cent_off":
            if line == "2250":
                return self.get_dca_file_name("flt")
            elif line == "1580":
                return self.get_dca_file_name("mfg")
            else:
                raise Exception("wrong line")

        elif part == "r2_cent_off":
            return self.get_dca_file_name("r2dw")
        else:
            raise Exception("unknown part")

    def build_df(self):

        for line in self.lines:

            for part, signal in self.signals[line].items():

                self.df = self.df.append({
                    "LINE": line,
                    "PART": part,
                    "DCAFILE": self.get_gauge_dca_file_name(line, part),
                    "SIGNAL": signal
                }, ignore_index=True)
