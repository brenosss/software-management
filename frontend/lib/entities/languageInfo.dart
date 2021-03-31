import 'package:flutter/material.dart';

class LanguageInfo {
  final String name;
  final String popularity;
  final String learning;
  final List<String> use;

  LanguageInfo(
      {@required this.name,
      @required this.popularity,
      @required this.learning,
      @required this.use});

  factory LanguageInfo.fromJson(Map<String, dynamic> json) {
    final data = LanguageInfo(
      name: json['Name'],
      popularity: json['Popularity'],
      learning: json['Learning'],
      use: json['Use'],
    );
    return data;
  }
}
