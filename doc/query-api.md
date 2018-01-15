# 12306 余票查询API

- 2017-01-21 API
```
curl -k "https://kyfw.12306.cO.train_date=2017-01-21&leftTicketDTO.from_station=GZQ&leftTicketDTO.to_station=WHN&purpose_codes=ADULT" && curl -k "https://kyfw.12306.cn/otn/leftTicket/queryA?leftTicketDTO.train_date=2017-01-21&leftTicketDTO.from_station=GZQ&leftTicketDTO.to_station=WHN&purpose_codes=ADULT"
```

```
{
    "queryLeftNewDTO": {
        "train_no": "240000G1010C",
        "station_train_code": "G101",
        "start_station_telecode": "VNP",
        "start_station_name": "北京南",
        "end_station_telecode": "AOH",
        "end_station_name": "上海虹桥",
        "from_station_telecode": "VNP",
        "from_station_name": "北京南",
        "to_station_telecode": "AOH",
        "to_station_name": "上海虹桥",
        "start_time": "06:44",
        "arrive_time": "12:38",
        "day_difference": "0",
        "train_class_name": "",
        "lishi": "05:54",
        "canWebBuy": "IS_TIME_NOT_BUY",
        "lishiValue": "354",
        "yp_info": "O055300032M0933000349174800012",
        "control_train_day": "20161229",
        "start_train_date": "20161201",
        "seat_feature": "O3M393",
        "yp_ex": "O0M090",
        "train_seat_feature": "3",
        "seat_types": "OM9",
        "location_code": "P2",
        "from_station_no": "01",
        "to_station_no": "11",
        "control_day": 59,
        "sale_time": "1230",
        "is_support_card": "1",
        "controlled_train_flag": "0",
        "controlled_train_message": "正常车次，不受控",
        "gg_num": "--",
        "gr_num": "--",
        "qt_num": "--",
        "rw_num": "--",
        "rz_num": "--",
        "tz_num": "--",
        "wz_num": "--",
        "yb_num": "--",
        "yw_num": "--",
        "yz_num": "--",
        "ze_num": "有",
        "zy_num": "有",
        "swz_num": "12"
    },
    "secretStr": "",
    "buttonTextInfo": "23:00-06:00系统维护时间"
}
```

- 2018-01-15 API
```
[ 预订 630000K2260L K226 GZQ LZJ GZQ WCN 20:20 09:24 13:04 N bwqeso67P92zuF963GRao0C3uigfAQVxWLwU2MBnA4Gh7maE8OHGyIQZVaw%3D 20180211 3 Q6 01 11 0 0    无   无  无 无     10401030 1413 0]
```