import pandas as pd
import numpy as np
import json

# generate part_table
df = pd.read_excel("part_table.xlsx")

part_dict = {}
part_dict["part_table"] = []
line_list = df["LINE"].unique()

for i, line in enumerate(line_list):
    part_dict["part_table"].append({})
    part_dict["part_table"][i]["line"] = int(line)
    part_dict["part_table"][i]["table"] = []
    table = part_dict["part_table"][i]["table"]

    records = df.loc[df["LINE"] == line]

    for j, idx in enumerate(records.index):
        table.append({})
        table[j]["part"] = records.loc[idx, "PART"]
        table[j]["dcafile"] = records.loc[idx, "DCAFILE"]
        # print(records.loc[idx, "SIGNAL"])
        table[j]["signal"] = (
            str(records.loc[idx, "SIGNAL"]).replace("\\\\", "\\"))

# print(part_dict)
with open("part_table.json", "w", encoding='utf-8') as jsfile:
    json.dump(part_dict, jsfile, indent=4, ensure_ascii=False)


def cast_dataframe(df, col):
    cast_tag = "cast_"
    if col.startswith(cast_tag):
        new_col = col.replace(cast_tag, "")
        std_list = [1, 2, 3, 4, 5, 6, 7]
        for std in std_list:
            df[new_col + str(std)] = df[col].apply(lambda x: x + str(std))
        df.drop([col], axis=1, inplace=True)
    else:
        pass


# generate task_table
df = pd.read_excel("factor_table.xlsx")
df = df.fillna("nan")

for col in df.columns:
    cast_dataframe(df, col)

print(df)


task_dict = {}
task_dict["factor_table"] = []
tbl_list = task_dict["factor_table"]

for i, col in enumerate(df.columns):
    tbl_list.append({})
    tbl_list[i]["seriesName"] = col
    tbl_list[i]["factorList"] = []
    task_list = tbl_list[i]["factorList"]
    for task in df[col]:
        if task.startswith("nan"):
            pass
        else:
            task_list.append(task)

with open("factor_table.json", "w", encoding='utf-8') as jsfile:
    json.dump(task_dict, jsfile, indent=4, ensure_ascii=False)
