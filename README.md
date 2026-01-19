### 2.3 Pipeline Sequence Diagram
This diagram illustrates the automated CI/CD flow, satisfying Requirement 2.3.

```mermaid
sequenceDiagram
    participant Dev as Developer
    participant GH as GitHub Repo
    participant GA as GitHub Actions
    participant Sonar as SonarCloud
    participant Docker as Docker Engine

    Note over Dev, Docker: Continuous Integration & Delivery Pipeline

    Dev->>GH: Push Code (git push)
    GH->>GA: Trigger Pipeline (Event: push)

    %% Job 1: Quality Control
    rect rgb(240, 248, 255)
        Note right of GA: Job 1: Quality Control
        GA->>GA: Checkout & Setup Go
        GA->>GA: Run Linter (Team Standards)
        GA->>GA: Run Unit Tests (Coverage)

        alt Branch is 'main'
            GA->>Sonar: Execute Static Analysis (Security/Bugs)
            Sonar-->>GA: Quality Gate Result
        else Branch is 'develop'
            Note right of GA: Skip Sonar (Resource Optimization)
        end
    end

    %% Job 2: Build & Release
    rect rgb(255, 240, 245)
        Note right of GA: Job 2: Build & Release (Needs Job 1)
        
        alt Branch is 'main'
            GA->>GA: Calculate SemVer (Conventional Commits)
            GA->>GH: Push Git Tag (e.g., v1.0.1)
            GA->>Docker: Build Image with Tag (app:v1.0.1)
        else Branch is 'develop'
            GA->>Docker: Build Image with SHA (app:a1b2c3d)
        end
        
        Docker-->>GA: Container Image Ready
    end
```

