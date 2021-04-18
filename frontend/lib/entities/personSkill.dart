import 'package:flutter/material.dart';
import 'package:frontend/entities/skill.dart';

class PersonSkill {
  final Skill skill;
  final int value;
  final int progress;

  PersonSkill(
      {@required this.value,
      @required this.progress,
      @required this.skill});

  factory PersonSkill.fromJson(Map<String, dynamic> json) {
    final data = PersonSkill(
      value: json['Value'],
      progress: json['Progress'],
      skill: Skill.fromJson(json['Skill'])

    );
    return data;
  }
}
