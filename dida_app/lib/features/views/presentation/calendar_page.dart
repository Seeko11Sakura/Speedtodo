import 'package:flutter/material.dart';

class CalendarPage extends StatefulWidget {
  const CalendarPage({super.key});

  @override
  State<CalendarPage> createState() => _CalendarPageState();
}

class _CalendarPageState extends State<CalendarPage> {
  DateTime focusedDay = DateTime.now();
  DateTime? selectedDay;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Calendar')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text('Calendar View', style: TextStyle(fontSize: 20)),
            const SizedBox(height: 16),
            Text(
              selectedDay != null
                  ? 'Selected: ${selectedDay!.toLocal()}'.split(' ')[0]
                  : 'No day selected',
            ),
          ],
        ),
      ),
    );
  }
}
