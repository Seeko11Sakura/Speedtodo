class FocusSession {
  const FocusSession({
    required this.id,
    required this.plannedMinutes,
    required this.status,
  });

  final int id;
  final int plannedMinutes;
  final String status;
}
