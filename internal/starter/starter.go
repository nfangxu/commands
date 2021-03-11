package starter

import (
    "context"
    "fmt"
    "github.com/nfangxu/commands/lib/base"
    "os"
    "path"
)

type Starter struct {
    Name string
    Resp string
}

func (s Starter) New(ctx context.Context, dir string) error {
    to := path.Join(dir, s.Name)
    if _, err := os.Stat(to); !os.IsNotExist(err) {
        return fmt.Errorf("%s already exists", s.Name)
    }
    fmt.Printf("Creating service %s\n", s.Name)
    repo := base.NewRepo(s.Resp)

    if err := repo.CopyTo(ctx, to, s.Name, []string{".git", ".github"}); err != nil {
        return err
    }
    os.Rename(
        path.Join(to, "cmd", "server"),
        path.Join(to, "cmd", s.Name),
    )
    return nil
}