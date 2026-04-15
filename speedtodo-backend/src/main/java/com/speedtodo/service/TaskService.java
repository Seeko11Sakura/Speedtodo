package com.speedtodo.service;
import com.speedtodo.model.Task;
import com.speedtodo.repository.TaskRepository;
import org.springframework.stereotype.Service;
import java.time.LocalDateTime;
import java.util.List;
@Service
public class TaskService {
    private final TaskRepository repository;
    public TaskService(TaskRepository repository) { this.repository = repository; }
    public List<Task> listTasks(String view) { return repository.findByStatus("active"); }
    public Task getById(Long id) { return repository.findById(id).orElseThrow(() -> new RuntimeException("Task not found")); }
    public Task create(Task task) { task.setStatus("active"); return repository.save(task); }
    public Task update(Long id, Task updated) {
        Task existing = getById(id);
        if (updated.getTitle() != null) existing.setTitle(updated.getTitle());
        if (updated.getNotes() != null) existing.setNotes(updated.getNotes());
        if (updated.getPriority() != 0) existing.setPriority(updated.getPriority());
        if (updated.getListId() != null) existing.setListId(updated.getListId());
        if (updated.getDueAt() != null) existing.setDueAt(updated.getDueAt());
        if (updated.getRemindAt() != null) existing.setRemindAt(updated.getRemindAt());
        existing.setAllDay(updated.isAllDay());
        if (updated.getRepeatRule() != null) existing.setRepeatRule(updated.getRepeatRule());
        return repository.save(existing);
    }
    public Task complete(Long id) { Task t = getById(id); t.setStatus("completed"); t.setCompletedAt(LocalDateTime.now()); return repository.save(t); }
    public Task reopen(Long id) { Task t = getById(id); t.setStatus("active"); t.setCompletedAt(null); return repository.save(t); }
    public void delete(Long id) { repository.deleteById(id); }
}