#include <stdio.h>
#include <math.h>

#include <gmp.h>
#include <mpfr.h>

typedef struct {
    mpfr_t x1, x2;
} V;

V getGradient(V p);
V mulV(V p, V grad, mpfr_t t);

int main (void)
{

    V p;
    mpfr_init2(p.x1,200);
    mpfr_init2(p.x2,200);
    mpfr_set_d (p.x1, 1.0, MPFR_RNDD);
    mpfr_set_d (p.x2, 1.0, MPFR_RNDD);

    mpfr_t t;
    mpfr_init2 (t, 200);
    mpfr_set_d (t, 0.01, MPFR_RNDD);



    unsigned int i;
    for (i = 1; i <= 1000000; i++)
    {
        V g = getGradient(p);
        p = mulV(p, g , t);
    }

    printf ("x1: ");
    mpfr_out_str (stdout, 10, 0, p.x1, MPFR_RNDD);
    printf ("      x2: ");
    mpfr_out_str (stdout, 10, 0, p.x2, MPFR_RNDD);
    putchar ('\n');
    mpfr_clear (p.x1);
    mpfr_clear (p.x2);
    mpfr_clear (t);
    mpfr_free_cache ();
    return 0;
}

V getGradient(V p) {
    V g;
    mpfr_init2(g.x1,200);
    mpfr_init2(g.x2,200);
    mpfr_set_d (g.x1, 0, MPFR_RNDD);
    mpfr_set_d (g.x2, 0, MPFR_RNDD);

    {
        mpfr_t v1, v2, v3, v4, a1,a2,a3;
        mpfr_init2(v1,200);
        mpfr_init2(v2,200);
        mpfr_init2(v3,200);
        mpfr_init2(v4,200);
        mpfr_init2(a1,200);
        mpfr_init2(a2,200);
        mpfr_init2(a3,200);

        mpfr_set_d (a1, M_E, MPFR_RNDD);
        mpfr_set_d (a2, M_E, MPFR_RNDD);
        mpfr_set_d (a3, M_E, MPFR_RNDD);

        mpfr_set_d (v1, 0, MPFR_RNDD);
        mpfr_set_d (v2, 3, MPFR_RNDD);
        mpfr_set_d (v3, -0.1, MPFR_RNDD);
        mpfr_set_d (v4, 0, MPFR_RNDD);

        mpfr_mul(v2, v2, p.x2, MPFR_RNDD);
        mpfr_add(v4, v2, v4, MPFR_RNDD);
        mpfr_add(v4, p.x1, v4, MPFR_RNDD);
        mpfr_add(v4, v3, v4, MPFR_RNDD);
        mpfr_exp(a1, v4, MPFR_RNDD);


        mpfr_set_d (v1, -3, MPFR_RNDD);
        mpfr_set_d (v2, -0.1, MPFR_RNDD);
        mpfr_set_d (v3, 0, MPFR_RNDD);
        //mpfr_set_d (v4, 0, MPFR_RNDD);

        mpfr_mul(v1, v1, p.x2, MPFR_RNDD);
        mpfr_add(v3, v1, v3, MPFR_RNDD);
        mpfr_add(v3, p.x1, v3, MPFR_RNDD);
        mpfr_add(v3, v2, v3, MPFR_RNDD);
        mpfr_exp(a2, v3, MPFR_RNDD);


        mpfr_set_d (v1, -0.1, MPFR_RNDD);
        mpfr_set_d (v2, 0, MPFR_RNDD);
        mpfr_set_d (v3, -1, MPFR_RNDD);

        mpfr_add(v2, v2, v1, MPFR_RNDD);
        mpfr_sub(v2, v2, p.x1, MPFR_RNDD);
        mpfr_exp(a3, v2, MPFR_RNDD);
        mpfr_mul(a3, v3, a3, MPFR_RNDD);

        mpfr_add(g.x1, a1, a2, MPFR_RNDD); 
        mpfr_add(g.x1, g.x1, a3, MPFR_RNDD); 

        mpfr_clear (v1);
        mpfr_clear (v2);
        mpfr_clear (v3);
        mpfr_clear (v4);
        mpfr_clear (a1);
        mpfr_clear (a2);
        mpfr_clear (a3);
    }

    {
        mpfr_t v1, v2, v3, v4, a1,a2,a3;
        mpfr_init2(v1,200);
        mpfr_init2(v2,200);
        mpfr_init2(v3,200);
        mpfr_init2(v4,200);
        mpfr_init2(a1,200);
        mpfr_init2(a2,200);
        mpfr_init2(a3,200);
        mpfr_set_d (v1, 3, MPFR_RNDD);
        mpfr_set_d (v2, -0.1, MPFR_RNDD);
        mpfr_set_d (v3, 3, MPFR_RNDD);
        mpfr_set_d (v4, 0, MPFR_RNDD);

        mpfr_set_d (a1, M_E, MPFR_RNDD);
        mpfr_set_d (a2, M_E, MPFR_RNDD);
        mpfr_set_d (a3, M_E, MPFR_RNDD);

        mpfr_mul(v1, v1, p.x2, MPFR_RNDD);
        mpfr_add(v4, v1, v4, MPFR_RNDD);
        mpfr_add(v4, p.x1, v4, MPFR_RNDD);
        mpfr_add(v4, v2, v4, MPFR_RNDD);
        mpfr_exp(a1, v4, MPFR_RNDD);
        mpfr_mul(a1, v3, a1, MPFR_RNDD);


        mpfr_set_d (v1, -3, MPFR_RNDD);
        mpfr_set_d (v2, -0.1, MPFR_RNDD);
        mpfr_set_d (v3, -3, MPFR_RNDD);
        mpfr_set_d (v4, 0, MPFR_RNDD);

        mpfr_mul(v1, v1, p.x2, MPFR_RNDD);
        mpfr_add(v4, v1, v4, MPFR_RNDD);
        mpfr_add(v4, p.x1, v4, MPFR_RNDD);
        mpfr_add(v4, v2, v4, MPFR_RNDD);
        mpfr_exp(a2, v4, MPFR_RNDD);
        mpfr_mul(a2, v3, a2, MPFR_RNDD);

        mpfr_add(g.x2, a1, a2, MPFR_RNDD); 

        mpfr_clear (v1);
        mpfr_clear (v2);
        mpfr_clear (v3);
        mpfr_clear (v4);
        mpfr_clear (a1);
        mpfr_clear (a2);
        mpfr_clear (a3);
    }



    return g;

}

V mulV(V p, V grad, mpfr_t t) {
    mpfr_t nX1, nX2;
    mpfr_init2(nX1,200);
    mpfr_init2(nX2,200);
    mpfr_mul(nX1, grad.x1, t, MPFR_RNDD);
    mpfr_mul(nX2, grad.x2, t, MPFR_RNDD);
    mpfr_sub(p.x1, p.x1, nX1, MPFR_RNDD);
    mpfr_sub(p.x2, p.x2, nX2, MPFR_RNDD);    
    mpfr_clear (nX1);
    mpfr_clear (nX2);
    mpfr_clear (grad.x1);
    mpfr_clear (grad.x2);
    return p;
}