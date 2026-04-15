package com.speedtodo.model;
import javax.persistence.*;
import java.time.LocalDateTime;
@Entity
public class FocusSession {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private Long taskId;
    private int plannedMinutes;
    private LocalDateTime startedAt;
    private LocalDateTime endedAt;
    private String status;
    private LocalDateTime createdAt;
    public FocusSession() {}
    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }
    public Long getTaskId() { return taskId; }
    public void setTaskId(Long taskId) { this.taskId = taskId; }
    public int getPlannedMinutes() { return plannedMinutes; }
    public void setPlannedMinutes(int plannedMinutes) { this.plannedMinutes = plannedMinutes; }
    public LocalDateTime getStartedAt() { return startedAt; }
    public void setStartedAt(LocalDateTime startedAt) { this.startedAt = startedAt; }
    public LocalDateTime getEndedAt() { return endedAt; }
    public void setEndedAt(LocalDateTime endedAt) { this.endedAt = endedAt; }
    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }
    public LocalDateTime getCreatedAt() { return createdAt; }
    public void setCreatedAt(LocalDateTime createdAt) { this.createdAt = createdAt; }
    @PrePersist protected void onCreate() { createdAt = LocalDateTime.now(); }
}