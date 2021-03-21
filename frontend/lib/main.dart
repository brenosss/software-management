import 'package:flutter/material.dart';
import 'package:frontend/pages/home.dart';
import 'package:frontend/pages/library.dart';

void main() {
  runApp(MaterialApp(
    initialRoute: '/',
    routes: {
      '/': (context) => HomeScreen(),
      '/library': (context) => LibraryScreen(),
    },
  ));
}
