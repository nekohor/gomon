from base_ponder import BasePonder


class RouphMillPonder(BasePonder):

    def __init__(self):

        super().__init__()

        self.dcafiles = {
            "2250": "R{}_POND_P{}",
            "1580": "R{}_POND_{}"
        }

        self.r1_signals = {}
        self.r1_signals["2250"] = {
            "r1_gap": "TN\\\\L_R1_R1GAP",
            "r1_leveling": "TN\\\\L_R1_R1LVGAP",
            "r1_roll_force": "TN\\\\L_R1_R1TRLFCE",
            "r1_diff_force": "TN\\\\L_R1_R1DIFRLFCE",
            "r1_top_torque": "TN\\\\L_R1_R1T_TFB",
            "r1_bot_torque": "TN\\\\L_R1_R1B_TFB",
            "r1_top_speed": "TN\\\\L_R1_R1T_SFB",
            "r1_bot_speed": "TN\\\\L_R1_R1B_SFB",
            "r1_bend_force": "TN\\\\L_R1_R1BENDFCE",
            "r1_ent_dsc_on": "TN\\\\L_R1_R1ENTDSCON",
            "r1_ext_dsc_on": "TN\\\\L_R1_R1EXTDSCON",
            "r1_trans_length": "TN\\\\L_R1_R1TRANSLEN"
        }
        self.r1_signals["1580"] = {
            "r1_trans_length": "TN\\\\L2_R1_R1TRANSLEN",
            "r1_gap": "TN\\\\L2_R1_R1GAP",
            "r1_leveling": "TN\\\\L2_R1_R1LVGAP",
            "r1_roll_force": "TN\\\\L2_R1_R1TRLFCE",
            "r1_diff_force": "TN\\\\L2_R1_R1DIFRLFCE",
            "r1_top_torque": "TN\\\\L2_R1_R1T_TFB",
            "r1_bot_torque": "TN\\\\L2_R1_R1B_TFB",
            "r1_top_speed": "TN\\\\L2_R1_R1T_SFB",
            "r1_bot_speed": "TN\\\\L2_R1_R1B_SFB",
            "r1_cur_pass": "TN\\\\L2_R1_INDPNOL2",
            "r1_total_pass": "TN\\\\L2_R1S_TPASSNO",
            "r1_ent_dsc_on": "TN\\\\L2_R1_R1ENTDSCON",
            "r1_ext_dsc_on": "TN\\\\L2_R1_R1EXTDSCON",
        }

        self.r2_signals = {}
        self.r2_signals["2250"] = {
            "r2_gap": "TN\\\\L_R2_R2GAP",
            "r2_leveling": "TN\\\\L_R2_R2LVGAP",
            "r2_roll_force": "TN\\\\L_R2_R2TRLFCE",
            "r2_diff_force": "TN\\\\L_R2_R2DIFRLFCE",
            "r2_top_torque": "TN\\\\L_R2_R2T_TFB",
            "r2_bot_torque": "TN\\\\L_R2_R2B_TFB",
            "r2_top_speed": "TN\\\\L_R2_R2T_SFB",
            "r2_bot_speed": "TN\\\\L_R2_R2B_SFB",
            "r2_bend_force": "TN\\\\L_R2_R2BENDFCE",
            "r2_ent_dsc_on": "TN\\\\L_R2_R2ENTDSCON",
            "r2_ext_dsc_on": "TN\\\\L_R2_R2EXTDSCON",
            "r2_trans_length": "TN\\\\L_R2_R2TRANSLEN",
        }
        self.r2_signals["1580"] = {
            "r2_trans_length": "TN\\\\L2_R2_R2TRANSLEN",
            "r2_gap": "TN\\\\L2_R2_R2GAP",
            "r2_leveling": "TN\\\\L2_R2_R2LVGAP",
            "r2_roll_force": "TN\\\\L2_R2_R2TRLFCE",
            "r2_diff_force": "TN\\\\L2_R2_R2DIFRLFCE",
            "r2_top_torque": "TN\\\\L2_R2_R2T_TFB",
            "r2_bot_torque": "TN\\\\L2_R2_R2B_TFB",
            "r2_top_speed": "TN\\\\L2_R2_R2T_SFB",
            "r2_bot_speed": "TN\\\\L2_R2_R2B_SFB",
            "r2_cur_pass": "TN\\\\L2_R2_INDPNOL2",
            "r2_total_pass": "TN\\\\L2_R2S_TPASSNO",
            "r2_ent_dsc_on": "TN\\\\L2_R2_R2ENTDSCON",
            "r2_ext_dsc_on": "TN\\\\L2_R2_R2EXTDSCON",
        }

            

    def build_df(self):

        for line in self.lines:
            for part, signal in self.r1_signals[line].items():
                for num in self.r1_nums:
                    self.df = self.df.append({
                        "LINE": line,
                        "PART": "{}{}".format(part, num),
                        "DCAFILE": self.dcafiles[line].format(1, num),
                        "SIGNAL": signal
                    }, ignore_index=True)

        for line in self.lines:
            for part, signal in self.r2_signals[line].items():
                for num in self.r2_nums:
                    self.df = self.df.append({
                        "LINE": line,
                        "PART": "{}{}".format(part, num),
                        "DCAFILE": self.dcafiles[line].format(2, num),
                        "SIGNAL": signal
                    }, ignore_index=True)
