import 'package:flutter/material.dart';

class Skill {
  final String name;
  final String type;

  Skill({@required this.name,
      @required this.type});

  factory Skill.fromJson(Map<String, dynamic> json) {
    final data = Skill(
      name: json['Name'],
      type: json['Type'],
    );
    return data;
  }
}