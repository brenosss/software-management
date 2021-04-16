import 'package:flutter/material.dart';

class Person {
  final String personId;
  final String name;
  final String email;
  final String phone;
  final String description;
  final String birthday;

  Person(
      {@required this.personId,
      @required this.name,
      @required this.email,
      @required this.phone,
      @required this.birthday,
      @required this.description});

  factory Person.fromJson(Map<String, dynamic> json) {
    final data = Person(
      personId: json['PersonId'],
      name: json['Name'],
      email: json['Email'],
      phone: json['Phone'],
      description: json['Description'],
      birthday: json['Birthday'],
    );
    return data;
  }
}
