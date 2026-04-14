import 'package:dida_app/features/views/presentation/calendar_page.dart';
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  testWidgets('shows calendar title', (tester) async {
    await tester.pumpWidget(const MaterialApp(home: CalendarPage()));
    expect(find.text('Calendar'), findsOneWidget);
  });
}
