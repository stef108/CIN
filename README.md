### 2.3 Pipeline Sequence Diagram
This diagram illustrates the automated CI/CD flow, satisfying Requirement 2.3.

```mermaid
sequenceDiagram
    participant Dev as Developer
    participant GH as GitHub Repo
    participant GA as GitHub Actions
    participant Sonar as SonarCloud
    participant Docker as Docker Engine
    participant Render as Render (Production)

    Note over Dev, Render: Continuous Integration & Delivery Pipeline

    Dev->>GH: Push Code (git push)
    GH->>GA: Trigger Pipeline (Event: push)

    %% Job 1: Quality Control
    rect rgb(240, 248, 255)
        Note right of GA: Job 1: Quality Control
        GA->>GA: Checkout & Setup Go
        GA->>GA: Run Linter (golangci-lint)
        GA->>GA: Run Unit Tests (Coverage)

        alt Branch is 'main'
            GA->>Sonar: Execute Static Analysis (Security)
            Sonar-->>GA: Quality Gate Result
        else Branch is 'develop'
            Note right of GA: Skip Sonar (Resource Optimization)
        end
    end

    %% Job 2: Build & Release (Packaging)
    rect rgb(255, 240, 245)
        Note right of GA: Job 2: Build & Release
        
        alt Branch is 'main'
            GA->>GA: Calculate SemVer (e.g., v1.0.0)
            GA->>GH: Push Git Tag
            GA->>Docker: Build Artifact with Tag
        else Branch is 'develop'
            GA->>Docker: Build Artifact with SHA
        end
    end

    %% Job 3: Deployment (The New Part)
    rect rgb(230, 255, 230)
        Note right of Render: Job 3: Deployment & Config
        
        alt Branch is 'main'
            GH->>Render: Trigger Auto-Deploy
            Note right of Render: Configuration Management:<br/>Inject APP_VERSION & PORT<br/>(Decoupled from Code)
            Render->>Render: Build Docker Container
            Render-->>Dev: Service Live (Health Check /)
        end
    end
```

