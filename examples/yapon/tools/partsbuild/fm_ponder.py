from base_ponder import BasePonder


class FinishingMillPonder(BasePonder):

    def __init__(self):

        super().__init__()

        # file tag with value "fx" or "fm" to choose which ponde file
        self.file_tag = "fx"

        self.stream_tags = {
            "<=4": 1,
            ">=5": 2,
        }

        self.signals = {}

        self.signals["2250"] = {

            # with stream tag
            "ct_gap": "TN\\\\L_AG{stream}_F{std}CTGAPFBK",
            "os_gap": "TN\\\\L_AG{stream}_F{std}OSGAPFBK",
            "ds_gap": "TN\\\\L_AG{stream}_F{std}DSGAPFBK",
            "gauge_thick": "TN\\\\L_AG{stream}_F{std}GMTHK",

            "total_roll_force": "TN\\\\L_F{std}H_TLRLFCFBK",
            "diff_roll_force": "TN\\\\L_F{std}H_DIFRLFCFBK",

            "os_bend_force": {
                "<=6": "TN\\\\L_F{std}LP_F{std}WBOSFCFBK",
                "==7": "TN\\\\L_F{std}B_F{std}WBOSFCFBK",
            },

            "ds_bend_force": {
                "<=6": "TN\\\\L_F{std}LP_F{std}WBDSFCFBK",
                "==7": "TN\\\\L_F{std}B_F{std}WBDSFCFBK",
            },

            "pos_shft": "TN\\\\L_CF_F{std}WRSTPOSFBK",
            "top_pos_shft": "TN\\\\L_CF_F{std}WRSTPOSFBK",
            "bot_pos_shft": "TN\\\\L_CF_F{std}WRSBPOSFBK",

            "wr_speed": "TN\\\\L_FM_F{std}_SFB",
            "mill_torque": "TN\\\\L_FM_F{std}_RTRQ",
            "wr_cool_flow": "TN\\\\L_FA_F{std}RLCFL",

            "rgs_on": "TN\\\\L_FA_F{std}RLGAPSPON",
            "top_rgs_flow": "TN\\\\L_FA_F{std}TRLGLUBFL",
            "bot_rgs_flow": "TN\\\\L_FA_F{std}BRLGLUBFL",


            "top_rgl_on": "TN\\\\L_FA_F{std}TRLGLUBOILON",
            "bot_rgl_on": "TN\\\\L_FA_F{std}BRLGLUBOILON",
            "top_rgl_flow": "TN\\\\L_FA_F{std}TRLGLUBFL",
            "bot_rgl_flow": "TN\\\\L_FA_F{std}BRLGLUBFL",

            "looper_torque": "TN\\\\L_LP_L{std}TRQFBK",
            "looper_force": "TN\\\\L_F{std}LP_FCFBK",
            "looper_angle": "TN\\\\L_F{std}LP_ANGFB",
            "looper_unit_tension": "TN\\\\L_F{std}LP_ANGFB",

            "strip_spray_on": "TN\\\\L_FA_F{std}STSPON",
            "bot_ent_strip_spray_on": "TN\\\\L_FA_F{std}BSTSPON",
            "bot_ext_strip_spray_on": "TN\\\\L_FA_F{std}BDSTSPON",

            "trans_length": "TN\\\\L_FM_F{std}TRANSLEN",

            "thermo_expan": "TN\\\\L_FTHS_F{std}WRTHEXP",
            "roll_wear": "TN\\\\L_FTHS_F{std}WRWEAR",
            "wear_expan": "TN\\\\L_FTHS_F{std}WRWTAMNT",

            "fume_sup_on": "TN\\\\L_FA_F{std}FUMESUPON",

            "isc_flow": "TN\\\\L_FA_F{cur_std}{nxt_std}ISCFL",
            "isc_flow_perc": "TN\\\\L_FA_F{cur_std}{nxt_std}ISCFLPER",
        }

        self.signals["1580"] = {

            # with stream tag
            "ct_gap": "TN\\\\L2_AG{stream}_F{std}CTGAPFBK",
            "os_gap": "TN\\\\L2_AG{stream}_F{std}OSGAPFBK",
            "ds_gap": "TN\\\\L2_AG{stream}_F{std}DSGAPFBK",
            "gauge_thick": "TN\\\\L2_AG{stream}_F{std}GMTHK",

            "strip_speed": "TN\\\\L2_FM_F{std}D_SPD",
            "trans_length": "TN\\\\L2_FM_F{std}TRANSLEN",
            "wr_speed": "TN\\\\L2_FM_F{std}_SFB",
            "lvl": "TN\\\\L2_AGC_F{std}LVGAP",

            "total_roll_force": "TN\\\\L2_F{std}H_TLRLFCFBK",
            "diff_roll_force": "TN\\\\L2_F{std}H_DIFRLFCFBK",

            "os_bend_force": {
                "<=6": "TN\\\\L2_F{std}H_F{std}WBOSFCFBK",
                "==7": "TN\\\\L2_F{std}H_F{std}WBOSFCFBK",
            },

            "ds_bend_force": {
                "<=6": "TN\\\\L2_F{std}H_F{std}WBDSFCFBK",
                "==7": "TN\\\\L2_F{std}H_F{std}WBDSFCFBK",
            },

            "bnd_frc": "TN\\\\L2_F{std}H_WBFCFBK",
            "mill_torque": "TN\\\\L2_FM_F{std}_RTRQ",

            "pos_shft": "TN\\\\L2_CF_F{std}WRSTPOSFBK",
            "top_pos_shft": "TN\\\\L2_CF_F{std}WRSTPOSFBK",
            "bot_pos_shft": "TN\\\\L2_CF_F{std}WRSBPOSFBK",

            "wr_cool_flow": "TN\\\\L2_FA_F{std}RLCFL",

            "top_rgl_on": "TN\\\\L2_FA_F{std}_TRGLON",
            "bot_rgl_on": "TN\\\\L2_FA_F{std}_BRGLON",
            "top_rgl_flow": "TN\\\\L2_FA_F{std}TRLGLUBFL",
            "bot_rgl_flow": "TN\\\\L2_FA_F{std}BRLGLUBFL",

            "looper_torque": "TN\\\\L2_LP_L{std}TRQFBK",
            "looper_force": "TN\\\\L2_F{std}H_L{std}FCFBK",
            "looper_angle": "TN\\\\L2_F{std}H_L{std}ANGFB",
            "looper_unit_tension": "TN\\\\L2_LP_L{std}TENFBK",

            "fume_sup_on": "TN\\\\L2_FA_F{std}_DSTSPSPON",
            "rgs_on": "TN\\\\L2_FA_F{std}_TANPSPON",
            "top_rgs_on": "TN\\\\L2_FA_F{std}_TANPSPON",
            "bot_rgs_on": "TN\\\\L2_FA_F{std}_BANPSPON",
            "strip_spray_on": "TN\\\\L2_FA_F{std}_STSPON",
            "wr_cool_flow_perc": "TN\\\\L2_FA_F{std}RLCFLPER",
            "ext_strip_spray_on": "TN\\\\L2_FA_F{std}D_STSPON",

            "roll_wear": "TN\\\\L2_FTHS_F{std}WRWEAR",
            "thermo_expan": "TN\\\\L2_FTHS_F{std}WRTHEXP",
            "manual_bend_interv": "TN\\\\L2_CF_F{std}MANBNFC",
            "wear_expan_wrbr": "TN\\\\L2_RTWMS_F{std}CWRBR",
            "wear_expan": "TN\\\\L2_RTWMS_F{std}CPCEWR",

            "isc_flow": "TN\\\\L2_FA_F{cur_std}{nxt_std}ISCFL",
            "isc_flow_perc": "TN\\\\L2_FA_F{cur_std}{nxt_std}ISCFLPER",
            "isc_trans_len": "TN\\\\L2_FA_F{cur_std}{nxt_std}ISC_TLEN",

        }

    def build_df(self):
        for line in self.lines:

            for part, signal in self.signals[line].items():

                for std in self.fm_stds:

                    self.append(line, part, signal, std)

    def append(self, line, part, signal, std):

        if part[-3:] == "gap" or part == "gauge_thick":

            signal_name = signal.format(
                stream=self.get_stream_tag(line, std),
                std=std)

        elif part[-10:] == "bend_force":

            if std == 7:
                signal_name = signal["==7"].format(std=std)
            else:
                signal_name = signal["<=6"].format(std=std)

        elif part[:3] == "isc":

            signal_name = signal.format(
                cur_std=std,
                nxt_std=std + 1)

        else:
            # print(line, part, signal, std)
            signal_name = signal.format(std=std)

        self.df = self.df.append({
            "LINE": line,
            "PART": "{}{}".format(part, std),
            "DCAFILE": self.get_fm_dca_file_name(std),
            "SIGNAL": signal_name
        }, ignore_index=True)

    def get_stream_tag(self, line, std):
        if line == "1580":
            return "C"
        elif line == "2250":
            if std <= 4:
                return "1"
            else:
                return "2"
        else:
            raise Exception("wrong line")

    def get_fm_dca_file_name(self, std):
        if self.file_tag == "fx":
            return self.get_dca_file_name("F{}".format(std))
        elif self.file_tag == "fm":
            return self.get_dca_file_name("FM")
        else:
            raise Exception("wrong file tag in FinishingMillPonder")

    def set_file_tag(self, tag):
        self.file_tag = tag
