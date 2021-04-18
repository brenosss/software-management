import 'package:flutter/material.dart';
import 'package:frontend/pages/home.dart';
import 'package:frontend/pages/library.dart';
import 'package:frontend/pages/personList.dart';
import 'package:frontend/pages/personDetail.dart';

void main() {
  runApp(MaterialApp(
    initialRoute: '/',
    routes: {
      '/': (context) => HomeScreen(),
      '/person/detail': (context) => PersonDetail(),
      '/person': (context) => PersonList(),
      '/library': (context) => LibraryScreen(),
    },
  ));
}
