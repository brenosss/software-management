import 'package:flutter/material.dart';
import 'package:frontend/entities/languageInfo.dart';

class LanguageInfoList extends StatelessWidget {
  final List<LanguageInfo> languageList;
  LanguageInfoList({this.languageList});

  @override
  Widget build(BuildContext context) {
    return GridView.count(
        padding: const EdgeInsets.all(50),
        crossAxisCount: 5,
        childAspectRatio: 2,
        children: languageList
            .map((languageItem) => LanguageInfoItem(languageItem: languageItem))
            .toList());
  }
}

class LanguageInfoItem extends StatelessWidget {
  final LanguageInfo languageItem;
  LanguageInfoItem({this.languageItem});

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Column(
        children: <Widget>[
          ListTile(
            leading: Icon(Icons.code),
            title: Text(languageItem.name),
          ),
          Text(languageItem.popularity, style: TextStyle(color: Colors.grey)),
          Text(languageItem.learning, style: TextStyle(color: Colors.grey)),
        ],
      ),
    );
  }
}
