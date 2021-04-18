import 'package:flutter/material.dart';
import 'package:frontend/entities/personSkill.dart';

class Person {
  final String personId;
  final String name;
  final String email;
  final String phone;
  final String description;
  final String birthday;
  List<PersonSkill> personSkills;

  Person(
      {@required this.personId,
      @required this.name,
      @required this.email,
      @required this.phone,
      @required this.birthday,
      @required this.description,
      this.personSkills});

  factory Person.fromJson(Map<String, dynamic> json) {
    List<PersonSkill> personSkills = [];
    if (json['PersonSkills'] != null) {
      personSkills = json['PersonSkills']
          .map<PersonSkill>((item) => PersonSkill.fromJson(item))
          .toList();
    }
    final data = Person(
        personId: json['PersonId'],
        name: json['Name'],
        email: json['Email'],
        phone: json['Phone'],
        description: json['Description'],
        birthday: json['Birthday'],
        personSkills: personSkills);
    return data;
  }
}
